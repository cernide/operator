package controllers

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"

	operationv1 "github.com/polyaxon/mloperator/api/v1"
	"github.com/polyaxon/mloperator/controllers/managers"
)

func (r *OperationReconciler) reconcileJobOp(ctx context.Context, instance *operationv1.Operation) (ctrl.Result, error) {
	// Reconcile the underlaying job
	return ctrl.Result{}, r.reconcileJob(ctx, instance)
}

func (r *OperationReconciler) reconcileJob(ctx context.Context, instance *operationv1.Operation) error {
	log := r.Log

	job := managers.GenerateJob(
		instance.Name,
		instance.Namespace,
		instance.Labels,
		instance.Annotations,
		instance.Termination.BackoffLimit,
		instance.Termination.ActiveDeadlineSeconds,
		instance.Termination.TTLSecondsAfterFinished,
		instance.BatchJobSpec.Template.Spec,
	)
	if err := ctrl.SetControllerReference(instance, job, r.Scheme); err != nil {
		log.V(1).Info("generateJob Error")
		return err
	}
	// Check if the Job already exists
	foundJob := &batchv1.Job{}
	justCreated := false
	err := r.Get(ctx, types.NamespacedName{Name: job.Name, Namespace: job.Namespace}, foundJob)
	if err != nil && apierrs.IsNotFound(err) {
		if instance.IsDone() {
			return nil
		}

		log.V(1).Info("Creating Job", "namespace", job.Namespace, "name", job.Name)
		err = r.Create(ctx, job)
		if err != nil {
			if updated := instance.LogWarning("OperatorCreateJob", err.Error()); updated {
				log.V(1).Info("Warning unable to create Job")
				if statusErr := r.Status().Update(ctx, instance); statusErr != nil {
					return statusErr
				}
				r.instanceSyncStatus(instance)
			}
			return err
		}
		justCreated = true
		instance.LogStarting()
		err = r.Status().Update(ctx, instance)
		r.instanceSyncStatus(instance)
	} else if err != nil {
		return err
	}
	// Update the job object and write the result back if there are any changes
	if !justCreated && !instance.IsDone() && managers.CopyJobFields(job, foundJob) {
		log.V(1).Info("Updating Job", "namespace", job.Namespace, "name", job.Name)
		err = r.Update(ctx, foundJob)
		if err != nil {
			return err
		}
	}

	// Check the job status
	if condUpdated := r.reconcileJobStatus(instance, *foundJob); condUpdated {
		log.V(1).Info("Reconciling Job status", "namespace", job.Namespace, "name", job.Name)
		err = r.Status().Update(ctx, instance)
		if err != nil {
			return err
		}
		r.instanceSyncStatus(instance)
	}

	return nil
}

func (r *OperationReconciler) reconcileJobStatus(instance *operationv1.Operation, job batchv1.Job) bool {
	now := metav1.Now()
	log := r.Log

	// Check the pods
	podStatus, reason, message := managers.HasUnschedulablePods(r.Client, instance)
	if podStatus == operationv1.OperationWarning {
		log.V(1).Info("Job has unschedulable pod(s)", "Reason", reason, "message", message)
		if updated := instance.LogWarning(reason, message); updated {
			log.V(1).Info("Job Logging Status Warning")
			return true
		}
		return false
	}

	if podStatus == operationv1.OperationStarting && job.Status.CompletionTime == nil {
		return false
	}

	if len(job.Status.Conditions) == 0 {
		if job.Status.Failed > 0 && job.Status.Active == 0 {
			if updated := instance.LogWarning("", ""); updated {
				log.V(1).Info("Job Logging Status Warning")
				return true
			}
		} else if job.Status.Active > 0 {
			if updated := instance.LogRunning(); updated {
				log.V(1).Info("Job Logging Status Running")
				return true
			}
		}
		return false
	}

	newJobCond := job.Status.Conditions[len(job.Status.Conditions)-1]

	if job.Status.Active == 0 && job.Status.Succeeded > 0 && managers.IsJobSucceeded(newJobCond) {
		if updated := instance.LogSucceeded(); updated {
			instance.Status.CompletionTime = &now
			log.V(1).Info("Job Logging Status Succeeded")
			return true
		}
	}

	if job.Status.CompletionTime != nil && job.Status.Succeeded > 0 && managers.IsJobSucceeded(newJobCond) {
		if updated := instance.LogSucceeded(); updated {
			instance.Status.CompletionTime = &now
			log.V(1).Info("Job Logging Status Succeeded with active non null")
			return true
		}
	}

	if job.Status.Failed > 0 && managers.IsJobFailed(newJobCond) {
		newMessage := operationv1.GetFailureMessage(newJobCond.Message, podStatus, reason, message)
		if updated := instance.LogFailed(newJobCond.Reason, newMessage); updated {
			instance.Status.CompletionTime = &now
			log.V(1).Info("Job Logging Status Failed", "Message", newMessage, "podStatus", podStatus, "PodMessage", message)
			return true
		}
	}
	return false
}

func (r *OperationReconciler) cleanUpJob(ctx context.Context, instance *operationv1.Operation) (ctrl.Result, error) {
	return r.handleTTL(ctx, instance)
}

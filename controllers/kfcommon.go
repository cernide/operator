package controllers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	operationv1 "github.com/polyaxon/operator/api/v1"
	"github.com/polyaxon/operator/controllers/kfapi"
	"github.com/polyaxon/operator/controllers/managers"
)

// Common logic for reconciling job status
func (r *OperationReconciler) reconcileKFJobStatus(instance *operationv1.Operation, job unstructured.Unstructured) (bool, error) {
	now := metav1.Now()
	log := r.Log

	// Check the pods
	podStatus, reason, message := managers.HasUnschedulablePods(r.Client, instance)
	if podStatus == operationv1.OperationWarning {
		log.V(1).Info("Service has unschedulable pod(s)", "Reason", reason, "message", message)
		if updated := instance.LogWarning(reason, message); updated {
			log.V(1).Info("Service Logging Status Warning")
			return true, nil
		}
		return false, nil
	}

	status, ok, unerr := unstructured.NestedFieldCopy(job.Object, "status")
	if !ok {
		if unerr != nil {
			log.Error(unerr, "NestedFieldCopy unstructured to status error")
			return false, nil
		}
		log.Info("NestedFieldCopy unstructured to status error",
			"err", "Status is not found in job")
		return false, nil
	}

	statusMap := status.(map[string]interface{})
	jobStatus := kfapi.JobStatus{}
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(statusMap, &jobStatus)
	if err != nil {
		log.Error(err, "Convert unstructured to status error")
		return false, err
	}

	if len(jobStatus.Conditions) == 0 {
		return false, nil
	}

	cond := jobStatus.Conditions[len(jobStatus.Conditions)-1]

	if cond.Type == kfapi.JobRunning || cond.Type == kfapi.JobCreated {
		instance.LogRunning()
		log.V(1).Info("Job Logging Status Running")
		return true, nil
	}

	if cond.Type == kfapi.JobSucceeded {
		instance.LogSucceeded()
		instance.Status.CompletionTime = &now
		log.V(1).Info("Job Logging Status Succeeded")
		return true, nil
	}

	if cond.Type == kfapi.JobFailed {
		newMessage := operationv1.GetFailureMessage(cond.Message, podStatus, reason, message)
		if updated := instance.LogFailed(cond.Reason, newMessage); updated {
			instance.Status.CompletionTime = &now
			log.V(1).Info("Job Logging Status Failed", "Message", newMessage, "podStatus", podStatus, "PodMessage", message)
			return true, nil
		}
	}

	if cond.Type == kfapi.JobRestarting {
		instance.LogWarning(cond.Reason, cond.Message)
		log.V(1).Info("Job Logging Status Warning")
		return true, nil
	}
	return false, nil
}

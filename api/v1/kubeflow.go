package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// This module was copied from kubeflow because
// some incompatible packages in the common repo

// CleanPodPolicy describes how to deal with pods when the job is finished.
// +k8s:openapi-gen=true
type CleanPodPolicy string

// Possible values for CleanPodPolicy
const (
	CleanPodPolicyUndefined CleanPodPolicy = ""
	CleanPodPolicyAll       CleanPodPolicy = "All"
	CleanPodPolicyRunning   CleanPodPolicy = "Running"
	CleanPodPolicyNone      CleanPodPolicy = "None"
)

// SchedulingPolicy encapsulates various scheduling policies of the distributed training
// job, for example `minAvailable` for gang-scheduling.
// +k8s:openapi-gen=true
type SchedulingPolicy struct {
	MinAvailable  *int32 `json:"minAvailable,omitempty"`
	Queue         string `json:"queue,omitempty"`
	PriorityClass string `json:"priorityClass,omitempty"`
}

// KFReplicaSpec is a description of kubeflow replica
// +k8s:openapi-gen=true
type KFReplicaSpec struct {
	// Replicas is the desired number of replicas of the given template.
	// If unspecified, defaults to 1.
	Replicas *int32 `json:"replicas,omitempty"`

	// Template is the object that describes the pod that
	// will be created for this replica. RestartPolicy in PodTemplateSpec
	// will be overide by RestartPolicy in ReplicaSpec
	Template corev1.PodTemplateSpec `json:"template,omitempty"`

	// Restart policy for all replicas within the job.
	// One of Always, OnFailure, Never and ExitCode.
	// Default to Never.
	RestartPolicy corev1.RestartPolicy `json:"restartPolicy,omitempty"`
}

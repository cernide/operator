//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	common "k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/polyaxon/operator/api/v1.BatchJobSpec":       schema_polyaxon_operator_api_v1_BatchJobSpec(ref),
		"github.com/polyaxon/operator/api/v1.MPIJobSpec":         schema_polyaxon_operator_api_v1_MPIJobSpec(ref),
		"github.com/polyaxon/operator/api/v1.Operation":          schema_polyaxon_operator_api_v1_Operation(ref),
		"github.com/polyaxon/operator/api/v1.OperationCondition": schema_polyaxon_operator_api_v1_OperationCondition(ref),
		"github.com/polyaxon/operator/api/v1.OperationStatus":    schema_polyaxon_operator_api_v1_OperationStatus(ref),
		"github.com/polyaxon/operator/api/v1.PytorchJobSpec":     schema_polyaxon_operator_api_v1_PytorchJobSpec(ref),
		"github.com/polyaxon/operator/api/v1.ServiceSpec":        schema_polyaxon_operator_api_v1_ServiceSpec(ref),
		"github.com/polyaxon/operator/api/v1.TFJobSpec":          schema_polyaxon_operator_api_v1_TFJobSpec(ref),
		"github.com/polyaxon/operator/api/v1.DaskJobSpec":        schema_polyaxon_operator_api_v1_DaskJobSpec(ref),
		"github.com/polyaxon/operator/api/v1.RayJobSpec":         schema_polyaxon_operator_api_v1_RayJobSpec(ref),
		"github.com/polyaxon/operator/api/v1.TerminationSpec":    schema_polyaxon_operator_api_v1_TerminationSpec(ref),
	}
}

func schema_polyaxon_operator_api_v1_BatchJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BatchJobSpec defines the desired state of a batch job",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"termination": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the number of retries before marking this job failed.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.TerminationSpec"),
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Description: "Template describes the pods that will be created.",
							Ref:         ref("k8s.io/api/core/v1.PodTemplateSpec"),
						},
					},
				},
				Required: []string{"template"},
			},
		},
		Dependencies: []string{
			"github.com/polyaxon/operator/api/v1.TerminationSpec", "k8s.io/api/core/v1.PodTemplateSpec"},
	}
}

func schema_polyaxon_operator_api_v1_MPIJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MPIJobSpec defines the desired state of an mpi job",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"termination": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the number of retries before marking this job failed.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.TerminationSpec"),
						},
					},
					"cleanPodPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Defines the policy for cleaning up pods after the Job completes. Defaults to Running.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"slotsPerWorker": {
						SchemaProps: spec.SchemaProps{
							Description: "CleanPodPolicy defines the policy that whether to kill pods after the job completes. Defaults to None.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"replicaSpecs": {
						SchemaProps: spec.SchemaProps{
							Description: "`MPIReplicaSpecs` contains maps from `MPIReplicaType` to `ReplicaSpec` that specify the MPI replicas to run.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kubeflow/tf-operator/pkg/apis/common/v1.ReplicaSpec"),
									},
								},
							},
						},
					},
				},
				Required: []string{"replicaSpecs"},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/tf-operator/pkg/apis/common/v1.ReplicaSpec", "github.com/polyaxon/operator/api/v1.TerminationSpec"},
	}
}

func schema_polyaxon_operator_api_v1_Operation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Operation is the Schema for the operations API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"batchJobSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of a job.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.BatchJobSpec"),
						},
					},
					"serviceSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of a Service.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.ServiceSpec"),
						},
					},
					"tfJobSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of a TFJob.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.TFJobSpec"),
						},
					},
					"pytorchJobSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of a PytorchJob.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.PytorchJobSpec"),
						},
					},
					"mpiJobSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of a MPIJob.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.MPIJobSpec"),
						},
					},
					"daskJobSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of a DaskJob.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.DaskJobSpec"),
						},
					},
					"rayJobSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of a RayJob.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.RayJobSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Current status of an op.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.OperationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/polyaxon/operator/api/v1.BatchJobSpec", "github.com/polyaxon/operator/api/v1.MPIJobSpec", "github.com/polyaxon/operator/api/v1.OperationStatus", "github.com/polyaxon/operator/api/v1.PytorchJobSpec", "github.com/polyaxon/operator/api/v1.ServiceSpec", "github.com/polyaxon/operator/api/v1.TFJobSpec", "github.com/polyaxon/operator/api/v1.RayJobSpec", "github.com/polyaxon/operator/api/v1.DaskJobSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_polyaxon_operator_api_v1_OperationCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OperationCondition defines the conditions of Operation or OpService",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is the type of the condition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastUpdateTime": {
						SchemaProps: spec.SchemaProps{
							Description: "The last time this condition was updated.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "The reason for this container condition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human readable message indicating details about the transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_polyaxon_operator_api_v1_OperationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OperationStatus defines the observed state of a job or a service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The latest available observations of an object's current state.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/polyaxon/operator/api/v1.OperationCondition"),
									},
								},
							},
						},
					},
					"startTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents the time when the job was acknowledged by the controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"completionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents the time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastReconcileTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents the last time when the job was reconciled. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/polyaxon/operator/api/v1.OperationCondition", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_polyaxon_operator_api_v1_PytorchJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PytorchJobSpec defines the desired state of a pytorch job",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"termination": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the number of retries before marking this job failed.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.TerminationSpec"),
						},
					},
					"cleanPodPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Defines the policy for cleaning up pods after the Job completes. Defaults to Running.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicaSpecs": {
						SchemaProps: spec.SchemaProps{
							Description: "A map of PyTorchReplicaType (type) to ReplicaSpec (value). Specifies the PyTorch cluster configuration. For example,\n  {\n    \"Master\": PyTorchReplicaSpec,\n    \"Worker\": PyTorchReplicaSpec,\n  }",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kubeflow/tf-operator/pkg/apis/common/v1.ReplicaSpec"),
									},
								},
							},
						},
					},
				},
				Required: []string{"replicaSpecs"},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/tf-operator/pkg/apis/common/v1.ReplicaSpec", "github.com/polyaxon/operator/api/v1.TerminationSpec"},
	}
}

func schema_polyaxon_operator_api_v1_ServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceSpec defines the desired state of a service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"termination": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the number of retries before marking this job failed.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.TerminationSpec"),
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"ports": {
						SchemaProps: spec.SchemaProps{
							Description: "optional List of ports to expose on the service",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"integer"},
										Format: "int32",
									},
								},
							},
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Description: "Template describes the pods that will be created.",
							Ref:         ref("k8s.io/api/core/v1.PodTemplateSpec"),
						},
					},
				},
				Required: []string{"template"},
			},
		},
		Dependencies: []string{
			"github.com/polyaxon/operator/api/v1.TerminationSpec", "k8s.io/api/core/v1.PodTemplateSpec"},
	}
}

func schema_polyaxon_operator_api_v1_TFJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TFJobSpec defines the desired state of a tf job",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"termination": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the number of retries before marking this job failed.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.TerminationSpec"),
						},
					},
					"cleanPodPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Defines the policy for cleaning up pods after the Job completes. Defaults to Running.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicaSpecs": {
						SchemaProps: spec.SchemaProps{
							Description: "A map of TFReplicaType (type) to ReplicaSpec (value). Specifies the TF cluster configuration. For example,\n  {\n    \"PS\": ReplicaSpec,\n    \"Worker\": ReplicaSpec,\n  }",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kubeflow/tf-operator/pkg/apis/common/v1.ReplicaSpec"),
									},
								},
							},
						},
					},
				},
				Required: []string{"replicaSpecs"},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/tf-operator/pkg/apis/common/v1.ReplicaSpec", "github.com/polyaxon/operator/api/v1.TerminationSpec"},
	}
}

func schema_polyaxon_operator_api_v1_RayJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RayJobSpec defines the desired state of a ray job",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"termination": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the number of retries before marking this job failed.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.TerminationSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/tf-operator/pkg/apis/common/v1.ReplicaSpec", "github.com/polyaxon/operator/api/v1.TerminationSpec"},
	}
}

func schema_polyaxon_operator_api_v1_DaskJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RayJobSpec defines the desired state of a ray job",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"termination": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the number of retries before marking this job failed.",
							Ref:         ref("github.com/polyaxon/operator/api/v1.TerminationSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/tf-operator/pkg/apis/common/v1.ReplicaSpec", "github.com/polyaxon/operator/api/v1.TerminationSpec"},
	}
}

func schema_polyaxon_operator_api_v1_TerminationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TerminationSpec defines the desired state of a job or a service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"backoffLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the number of retries before marking this job failed. Defaults to 1",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"activeDeadlineSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the duration (in seconds) since startTime during which the job can remain active before it is terminated. Must be a positive integer. This setting applies only to pods where restartPolicy is OnFailure or Always.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"ttlSecondsAfterFinished": {
						SchemaProps: spec.SchemaProps{
							Description: "ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
	}
}

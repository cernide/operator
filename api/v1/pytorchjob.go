package v1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2"
)

// PytorchJobSpec defines the desired state of a pytorch job
// +k8s:openapi-gen=true
type PytorchJobSpec struct {
	// Defines the policy for cleaning up pods after the Job completes.
	// Defaults to Running.
	CleanPodPolicy *CleanPodPolicy `json:"cleanPodPolicy,omitempty" protobuf:"bytes,1,opt,name=cleanPodPolicy"`

	// SchedulingPolicy defines the policy related to scheduling, e.g. gang-scheduling
	// +optional
	SchedulingPolicy *SchedulingPolicy `json:"schedulingPolicy,omitempty"  protobuf:"bytes,2,opt,name=schedulingPolicy"`

	// ElasticPolicy defines the policy related to elastic scaling
	ElasticPolicy *PytorchElasticPolicy `json:"elasticPolicy,omitempty"  protobuf:"bytes,3,opt,name=elasticPolicy"`

	// A map of ReplicaType (type) to ReplicaSpec (value). Specifies the PyTorch cluster configuration.
	// For example,
	//   {
	//     "Master": PyTorchReplicaSpec,
	//     "Worker": PyTorchReplicaSpec,
	//   }
	ReplicaSpecs map[PyTorchReplicaType]KFReplicaSpec `json:"replicaSpecs" protobuf:"bytes,4,opt,name=replicaSpecs"`

	// Number of workers per node; supported values: [auto, cpu, gpu, int].
	// For more, https://github.com/pytorch/pytorch/blob/26f7f470df64d90e092081e39507e4ac751f55d6/torch/distributed/run.py#L629-L658.
	// Defaults to auto.
	NprocPerNode *string `json:"nprocPerNode,omitempty" protobuf:"bytes,5,opt,name=replicaSpecs"`
}

// PyTorchReplicaType is the type for PyTorchReplica. Can be one of "Master" or "Worker".
type PyTorchReplicaType string

const (
	// PyTorchReplicaTypeMaster is the type of Master of distributed PyTorch
	PyTorchReplicaTypeMaster PyTorchReplicaType = "Master"

	// PyTorchReplicaTypeWorker is the type for workers of distributed PyTorch.
	PyTorchReplicaTypeWorker PyTorchReplicaType = "Worker"
)

type PytorchElasticPolicy struct {
	// minReplicas is the lower limit for the number of replicas to which the training job
	// can scale down.  It defaults to null.
	// +optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas, defaults to null.
	// +optional
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`

	RDZVBackend *RDZVBackend `json:"rdzvBackend,omitempty"`
	RDZVPort    *int32       `json:"rdzvPort,omitempty"`
	RDZVHost    *string      `json:"rdzvHost,omitempty"`
	RDZVID      *string      `json:"rdzvId,omitempty"`
	// RDZVConf contains additional rendezvous configuration (<key1>=<value1>,<key2>=<value2>,...).
	RDZVConf []RDZVConf `json:"rdzvConf,omitempty"`
	// Start a local standalone rendezvous backend that is represented by a C10d TCP store
	// on port 29400. Useful when launching single-node, multi-worker job. If specified
	// --rdzv_backend, --rdzv_endpoint, --rdzv_id are auto-assigned; any explicitly set values
	// are ignored.
	Standalone *bool `json:"standalone,omitempty"`
	// Number of workers per node; supported values: [auto, cpu, gpu, int].
	// Deprecated: This API is deprecated in v1.7+
	// Use .spec.nprocPerNode instead.
	NProcPerNode *int32 `json:"nProcPerNode,omitempty"`

	MaxRestarts *int32 `json:"maxRestarts,omitempty"`

	// Metrics contains the specifications which are used to calculate the
	// desired replica count (the maximum replica count across all metrics will
	// be used).  The desired replica count is calculated with multiplying the
	// ratio between the target value and the current value by the current
	// number of pods. Ergo, metrics used must decrease as the pod count is
	// increased, and vice-versa.  See the individual metric source types for
	// more information about how each type of metric must respond.
	// If not set, the HPA will not be created.
	// +optional
	Metrics []autoscalingv2.MetricSpec `json:"metrics,omitempty"`
}

type RDZVConf struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type RDZVBackend string

const (
	// BackendC10D is the rendezvous backend type for C10d.
	BackendC10D RDZVBackend = "c10d"
	// BackendETCD is the rendezvous backend type for ETCD.
	BackendETCD RDZVBackend = "etcd"
	// BackendETCDV2 is the rendezvous backend type for ETCD v2.
	BackendETCDV2 RDZVBackend = "etcd-v2"
)

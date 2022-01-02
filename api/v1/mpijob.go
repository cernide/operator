/*
Copyright 2018-2021 Polyaxon, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

// MPIJobSpec defines the desired state of an mpi job
// +k8s:openapi-gen=true
type MPIJobSpec struct {

	// Defines the policy for cleaning up pods after the Job completes.
	// Defaults to Running.
	CleanPodPolicy *CleanPodPolicy `json:"cleanPodPolicy,omitempty" protobuf:"bytes,1,opt,name=cleanPodPolicy"`

	// SchedulingPolicy defines the policy related to scheduling, e.g. gang-scheduling
	// +optional
	SchedulingPolicy *SchedulingPolicy `json:"schedulingPolicy,omitempty"  protobuf:"bytes,2,opt,name=schedulingPolicy"`

	// CleanPodPolicy defines the policy that whether to kill pods after the job completes.
	// Defaults to None.
	SlotsPerWorker *int32 `json:"slotsPerWorker,omitempty" protobuf:"bytes,3,opt,name=slotsPerWorker"`

	// SSHAuthMountPath is the directory where SSH keys are mounted.
	// Defaults to "/root/.ssh".
	SSHAuthMountPath string `json:"sshAuthMountPath,omitempty" protobuf:"bytes,4,opt,name=sshAuthMountPath"`

	// MPIImplementation is the MPI implementation.
	// Options are "OpenMPI" (default) and "Intel".
	Implementation MPIImplementation `json:"mpiImplementation,omitempty" protobuf:"bytes,5,opt,name=implementation"`

	// `MPIReplicaSpecs` contains maps from `MPIReplicaType` to `ReplicaSpec` that
	// specify the MPI replicas to run.
	ReplicaSpecs map[MPIReplicaType]KFReplicaSpec `json:"replicaSpecs" protobuf:"bytes,6,opt,name=replicaSpecs"`
}

// MPIReplicaType is the type for MPIReplica.
type MPIReplicaType string

const (
	// MPIReplicaTypeLauncher is the type for launcher replica.
	MPIReplicaTypeLauncher MPIReplicaType = "Launcher"

	// MPIReplicaTypeWorker is the type for worker replicas.
	MPIReplicaTypeWorker MPIReplicaType = "Worker"
)

type MPIImplementation string

const (
	MPIImplementationOpenMPI MPIImplementation = "OpenMPI"
	MPIImplementationIntel   MPIImplementation = "Intel"
)

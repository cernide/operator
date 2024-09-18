package managers

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	operationv1 "github.com/cernide/operator/api/v1"
)

// generateKFReplica generates a new ReplicaSpec
func generateKFReplica(replicaSpec operationv1.KFReplicaSpec, labels map[string]string, annotations map[string]string) *operationv1.KFReplicaSpec {
	l := make(map[string]string)
	for k, v := range replicaSpec.Template.GetLabels() {
		l[k] = v
	}
	for k, v := range labels {
		l[k] = v
	}
	a := make(map[string]string)
	for k, v := range replicaSpec.Template.GetAnnotations() {
		a[k] = v
	}
	for k, v := range annotations {
		a[k] = v
	}
	return &operationv1.KFReplicaSpec{
		Replicas:      replicaSpec.Replicas,
		RestartPolicy: replicaSpec.RestartPolicy,
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{Labels: l, Annotations: a},
			Spec:       replicaSpec.Template.Spec,
		},
	}
}

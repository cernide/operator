/*
Copyright 2018-2023 Polyaxon, Inc.

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

package controllers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	operationv1 "github.com/polyaxon/mloperator/api/v1"
)

// AddStartTime Adds starttime field by the reconciler
func (r *OperationReconciler) AddStartTime(ctx context.Context, instance *operationv1.Operation) error {
	if instance.Status.StartTime != nil {
		return nil
	}

	now := metav1.Now()
	log := r.Log

	log.V(1).Info("Setting StartTime", "Operation", instance.Name)
	instance.Status.StartTime = &now
	return r.Update(ctx, instance)
}

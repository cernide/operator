/*
Copyright 2018-2022 Polyaxon, Inc.

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

package plugins

import (
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-openapi/strfmt"

	httpRuntime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	netContext "golang.org/x/net/context"

	operationv1 "github.com/polyaxon/mloperator/api/v1"
	"github.com/polyaxon/mloperator/controllers/config"

	polyaxonSDK "github.com/polyaxon/sdks/go/http_client/v1/service_client"
	"github.com/polyaxon/sdks/go/http_client/v1/service_client/runs_v1"
	"github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

const (
	apiServerDefaultTimeout = 35 * time.Second

	// PolyaxonAuthToken polyaxon auth token
	PolyaxonAuthToken = "POLYAXON_AUTH_TOKEN"

	// PolyaxonInternalToken polyaxon internal token
	PolyaxonInternalToken = "POLYAXON_SECRET_INTERNAL_TOKEN"

	// PolyaxonAPIHost polyaxon api host
	PolyaxonAPIHost = "POLYAXON_PROXY_API_HOST"

	// PolyaxonAPIPort polyaxon api port
	PolyaxonAPIPort = "POLYAXON_PROXY_API_PORT"

	// PolyaxonStreamsHost polyaxon streams host
	PolyaxonStreamsHost = "POLYAXON_PROXY_STREAMS_HOST"

	// PolyaxonStreamsPort polyaxon api port
	PolyaxonStreamsPort = "POLYAXON_PROXY_STREAMS_PORT"
)

func polyaxonService(name string) string {
    if name == "InternalToken" {
        return "X-POLYAXON-INTERNAL"
    }
	return "X-POLYAXON-SERVICE"
}

func polyaxonAuth(name, value string) httpRuntime.ClientAuthInfoWriter {
	return httpRuntime.ClientAuthInfoWriterFunc(func(r httpRuntime.ClientRequest, _ strfmt.Registry) error {
		err := r.SetHeaderParam("Authorization", name+" "+value)
		if err != nil {
			return err
		}
		return r.SetHeaderParam(polyaxonService(name), "operator")
	})
}

func polyaxonHost(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// NotifyPolyaxonRunStatus creates polyaxon run status
func NotifyPolyaxonRunStatus(namespace, name, owner, project, uuid string, statusCond operationv1.OperationCondition, connections []string, log logr.Logger) error {
	token := config.GetStrEnv(PolyaxonAuthToken, "")
	host := polyaxonHost(config.GetStrEnv(PolyaxonStreamsHost, "localhost"), config.GetIntEnv(PolyaxonStreamsPort, 8000))

	plxClient := polyaxonSDK.New(httptransport.New(host, "", []string{"http"}), strfmt.Default)
	plxToken := polyaxonAuth("Token", token)

	ctx, cancel := netContext.WithTimeout(netContext.Background(), apiServerDefaultTimeout)
	defer cancel()

	params := &runs_v1.NotifyRunStatusParams{
		Namespace: namespace,
		Owner:     owner,
		Project:   project,
		UUID:      uuid,
		Body: &service_model.V1EntityNotificationBody{
			Name:        name,
			Connections: connections,
			Condition: &service_model.V1StatusCondition{
				LastTransitionTime: strfmt.DateTime(statusCond.LastTransitionTime.Time),
				LastUpdateTime:     strfmt.DateTime(statusCond.LastUpdateTime.Time),
				Message:            statusCond.Message,
				Reason:             statusCond.Reason,
				Status:             string(statusCond.Status),
				Type:               service_model.NewV1Statuses(service_model.V1Statuses(statusCond.Type)),
			},
		},
		Context: ctx,
	}
	_, _, err := plxClient.RunsV1.NotifyRunStatus(params, plxToken)
	if _, notFound := err.(*runs_v1.CollectRunLogsNotFound); notFound {
		return nil
	}
	return err
}

// LogPolyaxonRunStatus creates polyaxon run status
func LogPolyaxonRunStatus(owner, project, uuid string, statusCond operationv1.OperationCondition, log logr.Logger) error {
	token := config.GetStrEnv(PolyaxonAuthToken, "")
	port := config.GetIntEnv(PolyaxonAPIPort, 8000)
	scheme := "http"
	if port == 443 {
		scheme = "https"
	}
	host := polyaxonHost(config.GetStrEnv(PolyaxonAPIHost, "localhost"), port)

	plxClient := polyaxonSDK.New(httptransport.New(host, "", []string{scheme}), strfmt.Default)
	plxToken := polyaxonAuth("Token", token)

	ctx, cancel := netContext.WithTimeout(netContext.Background(), apiServerDefaultTimeout)
	defer cancel()

	params := &runs_v1.CreateRunStatusParams{
		Owner:   owner,
		Project: project,
		UUID:    uuid,
		Body: &service_model.V1EntityStatusBodyRequest{
			Condition: &service_model.V1StatusCondition{
				LastTransitionTime: strfmt.DateTime(statusCond.LastTransitionTime.Time),
				LastUpdateTime:     strfmt.DateTime(statusCond.LastUpdateTime.Time),
				Message:            statusCond.Message,
				Reason:             statusCond.Reason,
				Status:             string(statusCond.Status),
				Type:               service_model.NewV1Statuses(service_model.V1Statuses(statusCond.Type)),
			},
		},
		Context: ctx,
	}
	_, _, err := plxClient.RunsV1.CreateRunStatus(params, plxToken)
	if _, notFound := err.(*runs_v1.CreateRunStatusNotFound); notFound {
		log.Info("Operation create status; instance not found", "Project", project, "Instance", uuid)
		return nil
	}
	if _, forbidden := err.(*runs_v1.CreateRunStatusForbidden); forbidden {
		log.Info("Operation create status; forbidden", "Project", project, "Instance", uuid)
		return nil
	}
	if _, errorContent := err.(*runs_v1.CreateRunStatusDefault); errorContent {
		log.Info("Operation create status", "Error", errorContent, "Project", project, "Instance", uuid)
	}
	return err
}

// CollectPolyaxonRunLogs archives logs before removing the operation
func CollectPolyaxonRunLogs(namespace, owner, project, uuid string, kind string, log logr.Logger) error {
	token := config.GetStrEnv(PolyaxonInternalToken, "")
	host := polyaxonHost(config.GetStrEnv(PolyaxonStreamsHost, "localhost"), config.GetIntEnv(PolyaxonStreamsPort, 8000))

	plxClient := polyaxonSDK.New(httptransport.New(host, "", []string{"http"}), strfmt.Default)
	plxToken := polyaxonAuth("InternalToken", token)

	ctx, cancel := netContext.WithTimeout(netContext.Background(), apiServerDefaultTimeout)
	defer cancel()

	params := &runs_v1.CollectRunLogsParams{
		Namespace: namespace,
		Owner:     owner,
		Project:   project,
		UUID:      uuid,
		Kind:      kind,
		Context:   ctx,
	}
	_, _, err := plxClient.RunsV1.CollectRunLogs(params, plxToken)
	if _, notFound := err.(*runs_v1.CollectRunLogsNotFound); notFound {
		log.Info("Operation collect logs; instance not found", "Project", project, "Instance", uuid, "kind", kind)
		return nil
	}
	if _, forbidden := err.(*runs_v1.CollectRunLogsForbidden); forbidden {
		log.Info("Operation collect logs; forbidden", "Project", project, "Instance", uuid, "kind", kind)
		return nil
	}
	if _, errorContent := err.(*runs_v1.CollectRunLogsDefault); errorContent {
		log.Info("Operation collect logs", "Error", errorContent, "Project", project, "Instance", uuid, "kind", kind)
	}
	return err
}

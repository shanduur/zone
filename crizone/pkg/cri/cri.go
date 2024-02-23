package cri

import (
	"context"

	"github.com/shanduur/zone/cri-zone/pkg/cri/managers/containermanager"
	"github.com/shanduur/zone/cri-zone/pkg/cri/managers/containerstatsmanager"
	"github.com/shanduur/zone/cri-zone/pkg/cri/managers/imagemanagerservice"
	"github.com/shanduur/zone/cri-zone/pkg/cri/managers/podsandboxmanager"
	"github.com/shanduur/zone/cri-zone/pkg/cri/managers/runtimeversioner"
	cri "k8s.io/cri-api/pkg/apis"
	runtimev1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type RuntimeService struct {
	*containermanager.ContainerManager
	*containerstatsmanager.ContainerStatsManager
	*imagemanagerservice.ImageManagerService
	*podsandboxmanager.PodSandboxManager
	*runtimeversioner.RuntimeVersioner
}

var _ cri.RuntimeService = (*RuntimeService)(nil)

func (r *RuntimeService) UpdateRuntimeConfig(_ context.Context, _ *runtimev1.RuntimeConfig) error {
	return nil
}

func (r *RuntimeService) Status(_ context.Context, verbose bool) (*runtimev1.StatusResponse, error) {
	runtime := &runtimev1.RuntimeCondition{
		Type:   runtimev1.RuntimeReady,
		Status: true,
	}
	network := &runtimev1.RuntimeCondition{
		Type:   runtimev1.NetworkReady,
		Status: true,
	}

	resp := &runtimev1.StatusResponse{
		Status: &runtimev1.RuntimeStatus{
			Conditions: []*runtimev1.RuntimeCondition{
				runtime,
				network,
			},
		},
	}

	if verbose {
		resp.Info = map[string]string{}
	}

	return resp, nil
}

func (r *RuntimeService) RuntimeConfig(_ context.Context) (*runtimev1.RuntimeConfigResponse, error) {
	return &runtimev1.RuntimeConfigResponse{}, nil
}

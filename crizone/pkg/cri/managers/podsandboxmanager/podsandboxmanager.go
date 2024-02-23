package podsandboxmanager

import (
	"context"

	cri "k8s.io/cri-api/pkg/apis"
	runtimev1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type PodSandboxManager struct {
}

var _ cri.PodSandboxManager = (*PodSandboxManager)(nil)

func (m *PodSandboxManager) RunPodSandbox(ctx context.Context, config *runtimev1.PodSandboxConfig, runtimeHandler string) (string, error) {
	panic("unimplemented")
}

func (m *PodSandboxManager) StopPodSandbox(pctx context.Context, odSandboxID string) error {
	panic("unimplemented")
}

func (m *PodSandboxManager) RemovePodSandbox(ctx context.Context, podSandboxID string) error {
	panic("unimplemented")
}

func (m *PodSandboxManager) PodSandboxStatus(ctx context.Context, podSandboxID string, verbose bool) (*runtimev1.PodSandboxStatusResponse, error) {
	panic("unimplemented")
}

func (m *PodSandboxManager) ListPodSandbox(ctx context.Context, filter *runtimev1.PodSandboxFilter) ([]*runtimev1.PodSandbox, error) {
	panic("unimplemented")
}

func (m *PodSandboxManager) PortForward(context.Context, *runtimev1.PortForwardRequest) (*runtimev1.PortForwardResponse, error) {
	panic("unimplemented")
}

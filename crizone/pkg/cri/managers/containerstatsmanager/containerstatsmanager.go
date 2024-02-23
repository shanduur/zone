package containerstatsmanager

import (
	"context"

	cri "k8s.io/cri-api/pkg/apis"
	runtimev1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type ContainerStatsManager struct{}

var _ cri.ContainerStatsManager = (*ContainerStatsManager)(nil)

func (m *ContainerStatsManager) ContainerStats(ctx context.Context, containerID string) (*runtimev1.ContainerStats, error) {
	panic("unimplemented")
}

func (m *ContainerStatsManager) ListContainerStats(ctx context.Context, filter *runtimev1.ContainerStatsFilter) ([]*runtimev1.ContainerStats, error) {
	panic("unimplemented")
}

func (m *ContainerStatsManager) PodSandboxStats(ctx context.Context, podSandboxID string) (*runtimev1.PodSandboxStats, error) {
	panic("unimplemented")
}

func (m *ContainerStatsManager) ListPodSandboxStats(ctx context.Context, filter *runtimev1.PodSandboxStatsFilter) ([]*runtimev1.PodSandboxStats, error) {
	panic("unimplemented")
}

func (m *ContainerStatsManager) ListMetricDescriptors(ctx context.Context) ([]*runtimev1.MetricDescriptor, error) {
	panic("unimplemented")
}

func (m *ContainerStatsManager) ListPodSandboxMetrics(ctx context.Context) ([]*runtimev1.PodSandboxMetrics, error) {
	panic("unimplemented")
}

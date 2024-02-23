package containermanager

import (
	"context"
	"time"

	cri "k8s.io/cri-api/pkg/apis"
	runtimev1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type ContainerManager struct{}

var _ cri.ContainerManager = (*ContainerManager)(nil)

func (m *ContainerManager) CreateContainer(ctx context.Context, podSandboxID string, config *runtimev1.ContainerConfig, sandboxConfig *runtimev1.PodSandboxConfig) (string, error) {
	panic("unimplemented")
}

func (m *ContainerManager) StartContainer(ctx context.Context, containerID string) error {
	panic("unimplemented")
}

func (m *ContainerManager) StopContainer(ctx context.Context, containerID string, timeout int64) error {
	panic("unimplemented")
}

func (m *ContainerManager) RemoveContainer(ctx context.Context, containerID string) error {
	panic("unimplemented")
}

func (m *ContainerManager) ListContainers(ctx context.Context, filter *runtimev1.ContainerFilter) ([]*runtimev1.Container, error) {
	panic("unimplemented")
}

func (m *ContainerManager) ContainerStatus(ctx context.Context, containerID string, verbose bool) (*runtimev1.ContainerStatusResponse, error) {
	panic("unimplemented")
}

func (m *ContainerManager) UpdateContainerResources(ctx context.Context, containerID string, resources *runtimev1.ContainerResources) error {
	panic("unimplemented")
}

func (m *ContainerManager) ExecSync(ctx context.Context, containerID string, cmd []string, timeout time.Duration) (stdout []byte, stderr []byte, err error) {
	panic("unimplemented")
}

func (m *ContainerManager) Exec(context.Context, *runtimev1.ExecRequest) (*runtimev1.ExecResponse, error) {
	panic("unimplemented")
}

func (m *ContainerManager) Attach(ctx context.Context, req *runtimev1.AttachRequest) (*runtimev1.AttachResponse, error) {
	panic("unimplemented")
}

func (m *ContainerManager) ReopenContainerLog(ctx context.Context, ContainerID string) error {
	panic("unimplemented")
}

func (m *ContainerManager) CheckpointContainer(ctx context.Context, options *runtimev1.CheckpointContainerRequest) error {
	panic("unimplemented")
}

func (m *ContainerManager) GetContainerEvents(containerEventsCh chan *runtimev1.ContainerEventResponse) error {
	panic("unimplemented")
}

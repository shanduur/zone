package imagemanagerservice

import (
	"context"

	cri "k8s.io/cri-api/pkg/apis"
	runtimev1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type ImageManagerService struct{}

var _ cri.ImageManagerService = (*ImageManagerService)(nil)

func (m *ImageManagerService) ListImages(ctx context.Context, filter *runtimev1.ImageFilter) ([]*runtimev1.Image, error) {
	panic("unimplemented")
}

func (m *ImageManagerService) ImageStatus(ctx context.Context, image *runtimev1.ImageSpec, verbose bool) (*runtimev1.ImageStatusResponse, error) {
	panic("unimplemented")
}

func (m *ImageManagerService) PullImage(ctx context.Context, image *runtimev1.ImageSpec, auth *runtimev1.AuthConfig, podSandboxConfig *runtimev1.PodSandboxConfig) (string, error) {
	panic("unimplemented")
}

func (m *ImageManagerService) RemoveImage(ctx context.Context, image *runtimev1.ImageSpec) error {
	panic("unimplemented")
}

func (m *ImageManagerService) ImageFsInfo(ctx context.Context) (*runtimev1.ImageFsInfoResponse, error) {
	panic("unimplemented")
}

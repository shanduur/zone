package runtimeversioner

import (
	"context"

	"github.com/shanduur/zone/version"
	cri "k8s.io/cri-api/pkg/apis"
	runtimev1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type RuntimeVersioner struct {
	Name              string
	RuntimeVersion    string
	RuntimeAPIVersion string
}

var _ cri.RuntimeVersioner = (*RuntimeVersioner)(nil)

func (v *RuntimeVersioner) Version(_ context.Context, _ string) (*runtimev1.VersionResponse, error) {
	return &runtimev1.VersionResponse{
		Version:           version.KubeletVersion,
		RuntimeName:       version.Name,
		RuntimeVersion:    version.Version,
		RuntimeApiVersion: "v1",
	}, nil
}

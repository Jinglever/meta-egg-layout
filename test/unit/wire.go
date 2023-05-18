//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package unit

import (
	gomock "github.com/golang/mock/gomock"
	"meta-egg-layout/internal/biz"
	"meta-egg-layout/internal/common/resource"
	repo "meta-egg-layout/internal/repo"

	"github.com/google/wire"
)

func NewBizService(rsrc *resource.Resource, ctrl *gomock.Controller) *biz.BizService {
	panic(wire.Build(
		repo.MockProviderSet,
		biz.ProviderSet,
	))
}

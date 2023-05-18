//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package integrate

import (
	"meta-egg-layout/internal/biz"
	"meta-egg-layout/internal/common/resource"
	repo "meta-egg-layout/internal/repo"

	"github.com/google/wire"
)

func NewBizService(rsrc *resource.Resource) *biz.BizService {
	panic(wire.Build(
		repo.ProviderSet,
		biz.ProviderSet,
	))
}

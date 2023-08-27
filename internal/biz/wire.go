//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package biz

import (
	"meta-egg-layout/internal/common/resource"
	repo "meta-egg-layout/internal/repo"

	"github.com/google/wire"
)

func WireBizService(rsrc *resource.Resource) *BizService {
	panic(wire.Build(
		repo.ProviderSet,
		ProviderSet,
	))
}

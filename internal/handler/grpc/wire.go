//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package handler

import (
	"meta-egg-layout/internal/biz"
	"meta-egg-layout/internal/common/resource"
	repo "meta-egg-layout/internal/repo"

	"github.com/google/wire"
	//"meta-egg-layout/internal/usecase"
)

func WireHandler(rsrc *resource.Resource) *Handler {
	panic(wire.Build(
		repo.ProviderSet,
		biz.ProviderSet,
		//usecase.ProviderSet,
		ProviderSet,
	))
}

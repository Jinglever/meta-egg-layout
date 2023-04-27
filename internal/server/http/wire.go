//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package server

import (
	"meta-egg-layout/internal/biz"
	"meta-egg-layout/internal/common/resource"
	handler "meta-egg-layout/internal/handler/http"
	repo "meta-egg-layout/internal/repo"

	"github.com/google/wire"
	//"meta-egg-layout/internal/usecase"
)

func NewHandler(rsrc *resource.Resource) *handler.Handler {
	panic(wire.Build(
		repo.ProviderSet,
		biz.ProviderSet,
		//usecase.ProviderSet,
		handler.ProviderSet,
	))
}

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package server

import (
	"meta-egg-layout/internal/common/resource"
	domain "meta-egg-layout/internal/domain"
	handler "meta-egg-layout/internal/handler/http"
	repo "meta-egg-layout/internal/repo"

	"github.com/google/wire"
)

func NewHandler(rsrc *resource.Resource) *handler.Handler {
	panic(wire.Build(
		repo.ProviderSet,
		domain.ProviderSet,
		handler.ProviderSet,
	))
}

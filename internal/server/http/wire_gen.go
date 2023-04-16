// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"meta-egg-layout/internal/common/resource"
	"meta-egg-layout/internal/domain"
	"meta-egg-layout/internal/handler/http"
	"meta-egg-layout/internal/repo"
)

// Injectors from wire.go:

func NewHandler(rsrc *resource.Resource) *handler.Handler {
	genderRepo := repo.NewGenderRepo(rsrc)
	userRepo := repo.NewUserRepo(rsrc)
	domainUsecase := domain.NewDomainUsecase(rsrc, genderRepo, userRepo)
	handlerHandler := handler.NewHandler(rsrc, domainUsecase)
	return handlerHandler
}

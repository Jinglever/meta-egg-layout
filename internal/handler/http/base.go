/*
 * Generated by meta-egg.
 * WILL NOT be replace after re-generated. CAREFULLY EDIT.
 * Version: master-7741ad1-dirty
 * Author: meta-egg
 * Generated at: 2023-04-12 23:49
 */
package handler

import (
	gen "meta-egg-layout/gen/handler/http"
	"meta-egg-layout/internal/domain"

	"github.com/google/wire"
)

// ProviderSet is http handler providers.
var ProviderSet = wire.NewSet(
	NewHandler,
)

type Handler struct {
	gen.Handler

	// TODO: add your domain usecase
}

func NewHandler(
	domainUsecase *domain.DomainUsecase,

	// TODO: add domain usecase
) *Handler {
	return &Handler{
		Handler: gen.Handler{
			DomainUsecase: domainUsecase,
		},

		// TODO: setup domain usecase
	}
}

// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.1.0-EE
// Author: meta-egg
// Generated at: 2023-04-27 19:17

package handler

import (
	api "meta-egg-layout/api/meta_egg_layout"
	"meta-egg-layout/internal/biz"
	"meta-egg-layout/internal/common/resource"

	"github.com/google/wire"
)

// ProviderSet is grpc handler providers.
var ProviderSet = wire.NewSet(
	NewHandler,
)

var (
	DateLayout = "20060102"
	TimeLayout = "150405"
)

type Handler struct {
	api.UnimplementedMetaEggLayoutServer
	Resource   *resource.Resource
	BizService *biz.BizService
	// TODO: add your usecase
}

func NewHandler(
	rsrc *resource.Resource,
	bizService *biz.BizService,
	// TODO: add your usecase
) *Handler {
	return &Handler{
		Resource:   rsrc,
		BizService: bizService,
		// TODO: setup your usecase
	}
}

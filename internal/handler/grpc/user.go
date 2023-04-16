// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.0.0-EE
// Author: meta-egg
// Generated at: 2023-04-16 18:27
package handler

import (
	"context"
	api "meta-egg-layout/api/meta_egg_layout"

	"github.com/sirupsen/logrus"
)

// 获取用户详情
func (h *Handler) GetUserDetail(ctx context.Context, req *api.GetUserDetailRequest) (*api.GetUserDetailResponse, error) {
	mUser, err := h.DomainUsecase.GetUserByID(ctx, req.Id)
	if err != nil {
		logrus.Errorf("domainUsecase.GetUserByID failed, err: %v", err)
		return nil, err
	}
	return &api.GetUserDetailResponse{
		Id:        mUser.ID,
		Name:      mUser.Name,
		Gender:    mUser.Gender,
		CreatedBy: mUser.CreatedBy,
		CreatedAt: mUser.CreatedAt.Unix(),
		UpdatedBy: mUser.UpdatedBy,
		UpdatedAt: mUser.UpdatedAt.Unix(),
	}, nil
}

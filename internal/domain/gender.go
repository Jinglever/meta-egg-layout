// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.0.4-EE
// Author: meta-egg
// Generated at: 2023-04-22 00:40
package domain

import (
	"context"

	"meta-egg-layout/gen/model"
	"meta-egg-layout/internal/common/cerror"

	log "github.com/sirupsen/logrus"
)

func (u *DomainUsecase) GetGenderByID(ctx context.Context, id uint64) (*model.Gender, error) {
	mGender, err := u.GenderRepo.GetByID(ctx, id)
	if err != nil {
		log.Errorf("fail to get gender by id: %v", err)
		return nil, cerror.NotFound("gender not exist")
	}
	return mGender, nil
}

// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.2.1-IE
// Author: meta-egg
// Generated at: 2023-05-08 00:14

package biz

import (
	"context"

	"errors"
	"meta-egg-layout/gen/model"
	"meta-egg-layout/internal/common/cerror"

	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

func (u *BizService) GetGenderByID(ctx context.Context, id uint64) (*model.Gender, error) {
	mGender, err := u.GenderRepo.GetByID(ctx, id)
	if err != nil {
		log.Errorf("fail to get gender by id: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.NotFound("gender not exist")
		} else {
			return nil, cerror.Internal()
		}
	}
	return mGender, nil
}

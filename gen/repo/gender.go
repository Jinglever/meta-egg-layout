// Code generated by meta-egg. DO NOT EDIT.
// WILL BE replace after re-generated!
// DO NOT EDIT.
// Version: v1.4.1-IE
// Author: meta-egg
// Generated at: 2023-05-18 18:05

package repo

import (
	"context"

	"meta-egg-layout/gen/model"
	"meta-egg-layout/internal/common/resource"
	"meta-egg-layout/pkg/gormx"

	"gorm.io/gorm"
)

type GenderRepo interface {
	GetTX(ctx context.Context) *gorm.DB
	GetSematicByID(id uint64) string
	Gets(ctx context.Context, opts ...gormx.Option) ([]*model.Gender, error)
	GetByID(ctx context.Context, id uint64, opts ...gormx.Option) (*model.Gender, error)
	GetByIDs(ctx context.Context, ids []uint64, opts ...gormx.Option) ([]*model.Gender, error)
	Count(ctx context.Context, opts ...gormx.Option) (count int64, err error)
}

type GenderRepoImpl struct {
	Resource *resource.Resource
}

// get db
func (s *GenderRepoImpl) GetTX(ctx context.Context) *gorm.DB {
	// in case of transaction
	return s.Resource.DB.GetTX(ctx)
}

func (s *GenderRepoImpl) GetSematicByID(id uint64) string {
	switch id {
	case model.MetaGenderMale:
		return "MALE"
	case model.MetaGenderFemale:
		return "FEMALE"
	case model.MetaGenderUnknown:
		return "UNKNOWN"
	default:
		return ""
	}
}

func (s *GenderRepoImpl) Gets(ctx context.Context, opts ...gormx.Option) ([]*model.Gender, error) {
	var ms []*model.Gender
	tx := s.GetTX(ctx)
	for _, opt := range opts {
		tx = opt(tx)
	}
	if err := tx.Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

// get by primary key
func (s *GenderRepoImpl) GetByID(ctx context.Context, id uint64, opts ...gormx.Option) (*model.Gender, error) { //nolint
	var ms []*model.Gender
	opts = append(opts, gormx.Where(model.ColGenderID+" = ?", id), gormx.Limit(1))
	ms, err := s.Gets(ctx, opts...)
	if err != nil {
		return nil, err
	}
	if len(ms) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return ms[0], nil
}

// get by primary keys
func (s *GenderRepoImpl) GetByIDs(ctx context.Context, ids []uint64, opts ...gormx.Option) ([]*model.Gender, error) {
	if len(ids) == 0 {
		return make([]*model.Gender, 0), nil
	}
	var ms []*model.Gender
	opts = append(opts, gormx.Where(model.ColGenderID+" in (?)", ids))
	ms, err := s.Gets(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

// count
func (s *GenderRepoImpl) Count(ctx context.Context, opts ...gormx.Option) (count int64, err error) {
	tx := s.GetTX(ctx).Model(&model.Gender{})
	for _, opt := range opts {
		tx = opt(tx)
	}
	result := tx.Count(&count)
	if result.Error != nil {
		err = result.Error
	}
	return
}

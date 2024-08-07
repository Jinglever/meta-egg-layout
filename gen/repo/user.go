// Code generated by meta-egg. DO NOT EDIT.
// WILL BE replace after re-generated!
// DO NOT EDIT.
// Version: v1.4.1-IE
// Author: meta-egg
// Generated at: 2023-05-18 18:05

package repo

import (
	"context"
	"time"

	"meta-egg-layout/gen/model"
	"meta-egg-layout/internal/common/contexts"
	"meta-egg-layout/internal/common/resource"
	"meta-egg-layout/pkg/gormx"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetTX(ctx context.Context) *gorm.DB
	Gets(ctx context.Context, opts ...gormx.Option) ([]*model.User, error)
	GetByID(ctx context.Context, id uint64, opts ...gormx.Option) (*model.User, error)
	GetByIDs(ctx context.Context, ids []uint64, opts ...gormx.Option) ([]*model.User, error)
	Create(ctx context.Context, m *model.User) error
	CreateBatch(ctx context.Context, ms []*model.User) error
	Update(ctx context.Context, setCVs map[string]interface{}, incCVs map[string]interface{}, opts ...gormx.Option) (rowsAffected int64, err error)
	UpdateByID(ctx context.Context, id uint64, setCVs map[string]interface{}, incCVs map[string]interface{}) (rowsAffected int64, err error)
	UpdateByIDs(ctx context.Context, ids []uint64, setCVs map[string]interface{}, incCVs map[string]interface{}) (rowsAffected int64, err error)
	Delete(ctx context.Context, opts ...gormx.Option) (rowsAffected int64, err error)
	DeleteByID(ctx context.Context, id uint64) (rowsAffected int64, err error)
	DeleteByIDs(ctx context.Context, ids []uint64) (rowsAffected int64, err error)
	Count(ctx context.Context, opts ...gormx.Option) (count int64, err error)
}

type UserRepoImpl struct {
	Resource *resource.Resource
}

// get db
func (s *UserRepoImpl) GetTX(ctx context.Context) *gorm.DB {
	// in case of transaction
	return s.Resource.DB.GetTX(ctx)
}

func (s *UserRepoImpl) Gets(ctx context.Context, opts ...gormx.Option) ([]*model.User, error) {
	var ms []*model.User
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
func (s *UserRepoImpl) GetByID(ctx context.Context, id uint64, opts ...gormx.Option) (*model.User, error) { //nolint
	var ms []*model.User
	opts = append(opts, gormx.Where(model.ColUserID+" = ?", id), gormx.Limit(1))
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
func (s *UserRepoImpl) GetByIDs(ctx context.Context, ids []uint64, opts ...gormx.Option) ([]*model.User, error) {
	if len(ids) == 0 {
		return make([]*model.User, 0), nil
	}
	var ms []*model.User
	opts = append(opts, gormx.Where(model.ColUserID+" in (?)", ids))
	ms, err := s.Gets(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

// create single record
func (s *UserRepoImpl) Create(ctx context.Context, m *model.User) error {
	if me, ok := contexts.GetME(ctx); ok {
		meID := me.ID
		m.CreatedBy = &meID
		m.UpdatedBy = &meID
	}
	return s.GetTX(ctx).Create(m).Error
}

// create batch
func (s *UserRepoImpl) CreateBatch(ctx context.Context, ms []*model.User) error {
	if len(ms) == 0 {
		return nil
	}
	if me, ok := contexts.GetME(ctx); ok {
		for _, m := range ms {
			meID := me.ID
			m.CreatedBy = &meID
			m.UpdatedBy = &meID
		}
	}

	return s.GetTX(ctx).CreateInBatches(ms, CreateBatchNum).Error
}

func (s *UserRepoImpl) Update(ctx context.Context, setCVs map[string]interface{}, incCVs map[string]interface{}, opts ...gormx.Option) (rowsAffected int64, err error) {
	tx := s.GetTX(ctx).Model(&model.User{})
	for _, opt := range opts {
		tx = opt(tx)
	}
	if setCVs == nil {
		setCVs = make(map[string]interface{})
	}
	for k, v := range incCVs {
		setCVs[k] = gorm.Expr(k+" + ?", v)
	}
	if len(setCVs) == 0 {
		return 0, nil
	}
	if me, ok := contexts.GetME(ctx); ok {
		setCVs[model.ColUserUpdatedBy] = me.ID
	}
	result := tx.Updates(setCVs)
	if result.Error != nil {
		err = result.Error
	} else {
		rowsAffected = result.RowsAffected
	}
	return
}

// update by primary key
func (s *UserRepoImpl) UpdateByID(ctx context.Context, id uint64, setCVs map[string]interface{}, incCVs map[string]interface{}) (rowsAffected int64, err error) {
	return s.Update(ctx, setCVs, incCVs, gormx.Where(model.ColUserID+" = ?", id))
}

// update by primary keys
func (s *UserRepoImpl) UpdateByIDs(ctx context.Context, ids []uint64, setCVs map[string]interface{}, incCVs map[string]interface{}) (rowsAffected int64, err error) {
	if len(ids) == 0 {
		return 0, nil
	}
	return s.Update(ctx, setCVs, incCVs, gormx.Where(model.ColUserID+" in (?)", ids))
}

func (s *UserRepoImpl) Delete(ctx context.Context, opts ...gormx.Option) (rowsAffected int64, err error) {
	tx := s.GetTX(ctx).Model(&model.User{})
	for _, opt := range opts {
		tx = opt(tx)
	}
	var result *gorm.DB
	if me, ok := contexts.GetME(ctx); ok {
		result = tx.UpdateColumns(map[string]interface{}{
			model.ColUserDeletedBy: &(me.ID),
			model.ColUserDeletedAt: time.Now().Unix(),
		})
	} else {
		result = tx.Delete(&model.User{})
	}
	if result.Error != nil {
		err = result.Error
	} else {
		rowsAffected = result.RowsAffected
	}
	return
}

// delete by primary key
func (s *UserRepoImpl) DeleteByID(ctx context.Context, id uint64) (rowsAffected int64, err error) {
	return s.Delete(ctx, gormx.Where(model.ColUserID+" = ?", id))
}

// delete by primary keys
func (s *UserRepoImpl) DeleteByIDs(ctx context.Context, ids []uint64) (rowsAffected int64, err error) {
	if len(ids) == 0 {
		return 0, nil
	}
	return s.Delete(ctx, gormx.Where(model.ColUserID+" in (?)", ids))
}

// count
func (s *UserRepoImpl) Count(ctx context.Context, opts ...gormx.Option) (count int64, err error) {
	tx := s.GetTX(ctx).Model(&model.User{})
	for _, opt := range opts {
		tx = opt(tx)
	}
	result := tx.Count(&count)
	if result.Error != nil {
		err = result.Error
	}
	return
}

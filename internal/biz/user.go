// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v3.2.0-IE
// Author: meta-egg
// Generated at: 2024-01-01 22:48

package biz

import (
	"context"
	"time"

	"errors"
	"meta-egg-layout/gen/model"
	"meta-egg-layout/internal/common/cerror"
	"meta-egg-layout/internal/common/contexts"
	"meta-egg-layout/internal/repo/option"

	"gorm.io/gorm"

	jgstr "github.com/Jinglever/go-string"
)

type UserBO struct {
	ID        uint64     `json:"id"`         //
	Name      *string    `json:"name"`       // 用户名
	Gender    uint64     `json:"gender"`     // 性别
	Age       uint8      `json:"age"`        // 年龄
	IsOnJob   bool       `json:"is_on_job"`  // 是否在职
	Birthday  *time.Time `json:"birthday"`   // 生日
	CreatedBy *uint64    `json:"created_by"` // 创建者
	CreatedAt time.Time  `json:"created_at"` // 创建时间
	UpdatedBy *uint64    `json:"updated_by"` // 更新者
	UpdatedAt time.Time  `json:"updated_at"` // 更新时间
}

func (b *BizService) ToUserBO(ctx context.Context, m *model.User) (*UserBO, error) {
	return &UserBO{
		ID:        m.ID,
		Name:      m.Name,
		Gender:    m.Gender,
		Age:       m.Age,
		IsOnJob:   m.IsOnJob,
		Birthday:  m.Birthday,
		CreatedBy: m.CreatedBy,
		CreatedAt: m.CreatedAt,
		UpdatedBy: m.UpdatedBy,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func (b *BizService) CreateUser(ctx context.Context, obj *UserBO) error {
	log := contexts.GetLogger(ctx).
		WithField("obj", jgstr.JsonEncode(obj))
	m := &model.User{
		Name:     obj.Name,
		Gender:   obj.Gender,
		Age:      obj.Age,
		IsOnJob:  obj.IsOnJob,
		Birthday: obj.Birthday,
	}
	return b.Resource.DB.Transaction(ctx, func(txCtx context.Context) error {
		err := b.UserRepo.Create(txCtx, m)
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				log.WithError(err).Error("fail to create user, duplicated key")
				return cerror.AlreadyExists(err.Error())
			}

			log.WithError(err).Error("fail to create user")
			return cerror.Internal(err.Error())
		}
		bo, err := b.ToUserBO(txCtx, m)
		if err != nil {
			log.WithError(err).Error("fail to convert user model to UserBO")
			return cerror.Internal(err.Error())
		}
		*obj = *bo
		return nil
	})
}

func (b *BizService) GetUserByID(ctx context.Context, id uint64) (*UserBO, error) {
	log := contexts.GetLogger(ctx).
		WithField("id", id)
	mUser, err := b.UserRepo.GetByID(ctx, id)
	if err != nil {
		log.WithError(err).Error("fail to get user by id")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.NotFound(err.Error())
		} else {
			return nil, cerror.Internal(err.Error())
		}
	}
	userBO, err := b.ToUserBO(ctx, mUser)
	if err != nil {
		log.WithError(err).Error("fail to convert user model to UserBO")
		return nil, cerror.Internal(err.Error())
	}
	return userBO, nil
}

type UserListBO struct {
	ID     uint64  `json:"id"`     //
	Name   *string `json:"name"`   // 用户名
	Gender uint64  `json:"gender"` // 性别
}

func (b *BizService) ToUserListBO(ctx context.Context, ms []*model.User) ([]*UserListBO, error) {
	list := make([]*UserListBO, 0, len(ms))
	for i := range ms {
		list = append(list, &UserListBO{
			ID:     ms[i].ID,
			Name:   ms[i].Name,
			Gender: ms[i].Gender,
		})
	}
	return list, nil
}

type UserFilterOption struct {
	Gender  *uint64 // 性别
	IsOnJob *bool   // 是否在职
}

type UserListOption struct {
	Pagination *option.PaginationOption
	Order      *option.OrderOption
	Filter     *UserFilterOption
}

func (b *BizService) GetUserList(ctx context.Context, opt *UserListOption) ([]*UserListBO, int64, error) {
	log := contexts.GetLogger(ctx).
		WithField("opt", jgstr.JsonEncode(opt))
	msUser, total, err := b.UserRepo.GetList(ctx, &option.UserListOption{
		Pagination: opt.Pagination,
		Order:      opt.Order,
		Filter: &option.UserFilterOption{
			Gender:  opt.Filter.Gender,
			IsOnJob: opt.Filter.IsOnJob,
		},
		Select: []interface{}{
			model.ColUserID,
			model.ColUserName,
			model.ColUserGender,
		},
	})
	if err != nil {
		log.WithError(err).Error("fail to get user list")
		return nil, 0, cerror.Internal(err.Error())
	}
	list, err := b.ToUserListBO(ctx, msUser)
	if err != nil {
		log.WithError(err).Error("fail to convert user model to UserListBO")
		return nil, 0, cerror.Internal(err.Error())
	}
	return list, total, nil
}

type UserSetOption struct {
	Name     *string
	Gender   *uint64
	Age      *uint8
	IsOnJob  *bool
	Birthday *time.Time
}

func (b *BizService) UpdateUserByID(ctx context.Context, id uint64, setOpt *UserSetOption) error {
	log := contexts.GetLogger(ctx).
		WithField("id", id).
		WithField("setOpt", jgstr.JsonEncode(setOpt))
	// assemble setCVs
	setCVs := make(map[string]interface{})
	if setOpt != nil {
		if setOpt.Name != nil {
			setCVs[model.ColUserName] = *setOpt.Name
		}
		if setOpt.Gender != nil {
			setCVs[model.ColUserGender] = *setOpt.Gender
		}
		if setOpt.Age != nil {
			setCVs[model.ColUserAge] = *setOpt.Age
		}
		if setOpt.IsOnJob != nil {
			setCVs[model.ColUserIsOnJob] = *setOpt.IsOnJob
		}
		if setOpt.Birthday != nil {
			setCVs[model.ColUserBirthday] = *setOpt.Birthday
		}
	}

	if len(setCVs) == 0 {
		return nil
	}
	_, err := b.UserRepo.UpdateByID(ctx, id, setCVs, nil)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.WithError(err).Error("fail to update user, duplicated key")
			return cerror.AlreadyExists(err.Error())
		}

		log.WithError(err).Error("fail to update user")
		return cerror.Internal(err.Error())
	}
	return nil
}

func (b *BizService) DeleteUserByID(ctx context.Context, id uint64) error {
	log := contexts.GetLogger(ctx).
		WithField("id", id)
	_, err := b.UserRepo.DeleteByID(ctx, id)
	if err != nil {
		log.WithError(err).Error("fail to delete user")
		return cerror.Internal(err.Error())
	}
	return nil
}

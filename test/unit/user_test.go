package unit

import (
	"context"
	"meta-egg-layout/gen/model"
	"testing"

	repomock "meta-egg-layout/internal/repo/mock"
	gormxmock "meta-egg-layout/pkg/gormx/mock"

	jgptr "github.com/Jinglever/go-pointer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc := NewBizService(GetResource(ctrl), ctrl)
	m1 := svc.Resource.DB.(*gormxmock.MockDB)
	m1.EXPECT().
		Transaction(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
			return fn(ctx)
		}).AnyTimes()
	m2 := svc.UserRepo.(*repomock.MockUserRepo)
	m2.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, m *model.User) error {
			m.ID = 1
			return nil
		})
	ctx := context.TODO()
	mUser := &model.User{
		Name:    jgptr.NewString("test"),
		Gender:  1,
		Age:     32,
		IsOnJob: true,
	}
	err := svc.Resource.DB.Transaction(ctx, func(ctx context.Context) error {
		err := svc.CreateUser(ctx, mUser)
		assert.NoError(t, err)
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), mUser.ID)
}

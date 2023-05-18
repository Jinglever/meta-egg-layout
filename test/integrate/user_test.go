package integrate

import (
	"context"
	"meta-egg-layout/gen/model"
	"meta-egg-layout/internal/biz"
	"testing"

	jgptr "github.com/Jinglever/go-pointer"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	svc := NewBizService(GetResource())
	ctx := context.TODO()
	err := svc.CreateUser(ctx, &model.User{
		Name:    jgptr.NewString("test"),
		Gender:  1,
		Age:     32,
		IsOnJob: true,
	})
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	svc := NewBizService(GetResource())
	ctx := context.TODO()
	user, err := svc.GetUserByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "test", *user.Name)
}

func TestUpdateUser(t *testing.T) {
	svc := NewBizService(GetResource())
	ctx := context.TODO()
	err := svc.UpdateUserByID(ctx, 1, &biz.UserUpdateOption{
		Set: &biz.UserSetOption{
			Age: jgptr.NewUint8(33),
		},
	})
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	svc := NewBizService(GetResource())
	ctx := context.TODO()
	err := svc.DeleteUserByID(ctx, 1)
	assert.NoError(t, err)
}

func TestCreateUserTX(t *testing.T) {
	svc := NewBizService(GetResource())
	ctx := context.TODO()
	err := svc.Resource.DB.Transaction(ctx, func(txCtx context.Context) error {
		err := svc.CreateUser(txCtx, &model.User{
			Name:    jgptr.NewString("test"),
			Gender:  1,
			Age:     32,
			IsOnJob: true,
		})
		assert.NoError(t, err)
		return nil
	})
	assert.NoError(t, err)
}

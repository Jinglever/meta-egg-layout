package integrate

import (
	"context"
	"fmt"
	"meta-egg-layout/internal/biz"
	"testing"

	jgptr "github.com/Jinglever/go-pointer"
	jgstr "github.com/Jinglever/go-string"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	svc := biz.WireBizService(GetResource())
	ctx := context.TODO()
	err := svc.CreateUser(ctx, &biz.UserBO{
		Name:    jgptr.NewString("test"),
		Gender:  1,
		Age:     32,
		IsOnJob: true,
	})
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	svc := biz.WireBizService(GetResource())
	ctx := context.TODO()
	user, err := svc.GetUserByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "test", *user.Name)
}

func TestGetUsers(t *testing.T) {
	svc := biz.WireBizService(GetResource())
	ctx := context.TODO()
	users, _, err := svc.GetUserList(ctx, &biz.UserListOption{
		Filter: &biz.UserFilterOption{
			IsOnJob: jgptr.NewBool(true),
		},
	})
	assert.NoError(t, err)
	fmt.Println(jgstr.JsonEncode(users))
}

func TestUpdateUser(t *testing.T) {
	svc := biz.WireBizService(GetResource())
	ctx := context.TODO()
	err := svc.UpdateUserByID(ctx, 1, &biz.UserSetOption{
		Age:     jgptr.NewUint8(33),
		IsOnJob: jgptr.NewBool(true),
	})
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	svc := biz.WireBizService(GetResource())
	ctx := context.TODO()
	err := svc.DeleteUserByID(ctx, 1)
	assert.NoError(t, err)
}

func TestCreateUserTX(t *testing.T) {
	svc := biz.WireBizService(GetResource())
	ctx := context.TODO()
	err := svc.Resource.DB.Transaction(ctx, func(txCtx context.Context) error {
		err := svc.CreateUser(txCtx, &biz.UserBO{
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

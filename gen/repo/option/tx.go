// Code generated by meta-egg. DO NOT EDIT.
// WILL BE replace after re-generated!
// DO NOT EDIT.
// Version: v1.1.0-EE
// Author: meta-egg
// Generated at: 2023-04-27 19:17

package option

import (
	"context"

	"gorm.io/gorm"
)

type CtxKeyTX struct{}

func SetTX(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, CtxKeyTX{}, tx)
}

func GetTX(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(CtxKeyTX{}).(*gorm.DB)
	return tx, ok
}

func DelTX(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxKeyTX{}, nil)
}
/*
 * Generated by meta-egg.
 * Might be replace after re-generated. DO NOT EDIT!
 * Version: v0.9.0-96bf0e0
 * Author: meta-egg
 * Generated at: 2023-04-13 09:27
 */
package domain

import (
	repo "meta-egg-layout/internal/repo"

	"github.com/google/wire"
)

// ProviderSet is domain providers.
var ProviderSet = wire.NewSet(
	NewDomainUsecase,
)

// 匿名用例
type DomainUsecase struct {
	GenderRepo repo.GenderRepo
	UserRepo   repo.UserRepo
}

func NewDomainUsecase(
	genderRepo repo.GenderRepo,
	userRepo repo.UserRepo,

) *DomainUsecase {
	return &DomainUsecase{
		GenderRepo: genderRepo,
		UserRepo:   userRepo,
	}
}

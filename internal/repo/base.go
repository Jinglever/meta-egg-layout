// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.4.1-IE
// Author: meta-egg
// Generated at: 2023-05-18 22:30

package repo

import (
	mock "meta-egg-layout/internal/repo/mock"

	"github.com/google/wire"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(
	NewUserRepo,
	NewGenderRepo,
)

// MockProviderSet is mock repo providers.
var MockProviderSet = wire.NewSet(
	mock.NewMockUserRepo,
	wire.Bind(new(UserRepo), new(*mock.MockUserRepo)),
	mock.NewMockGenderRepo,
	wire.Bind(new(GenderRepo), new(*mock.MockGenderRepo)),
)

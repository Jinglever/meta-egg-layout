// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.0.19-EE
// Author: meta-egg
// Generated at: 2023-04-27 11:37
package domain

import (
	crud "meta-egg-layout/internal/domain/crud"

	"github.com/google/wire"
)

// ProviderSet is domain providers.
var ProviderSet = wire.NewSet(
	crud.NewCrudUsecase,
)

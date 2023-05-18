package unit

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	rsrc := GetResource(ctrl)
	assert.NotNil(t, rsrc)
}

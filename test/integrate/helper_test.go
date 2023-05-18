package integrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResource(t *testing.T) {
	rsrc := GetResource()
	assert.NotNil(t, rsrc)
}

package graphrest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResource(t *testing.T) {
	re := NewResource("/foo/123?fields=f1,f2")
	assert.Equal(t, "foo", re.Name)
	assert.Equal(t, "123", re.ID)

	re = NewResource("/foo/123/bar?fields=f1,f2")
	assert.Equal(t, "foo", re.Name)
	assert.Equal(t, "123", re.ID)
}

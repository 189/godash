package godash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWhole(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	v := IsWhole(3.1)
	is.False(v)

	v = IsWhole(3)
	is.True(v)
}

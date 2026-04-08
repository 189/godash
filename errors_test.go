package godash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	is := assert.New(t)

	b := Try(func() error {
		panic("ho no, don't panic here")
	})

	is.False(b)

	b = Try(func() error {
		return nil
	})
	is.True(b)
}

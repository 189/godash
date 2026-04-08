package godash

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsBool(t *testing.T) {
	is := assert.New(t)
	is.Equal(IsBool(true), true)
}

func TestIsChannel(t *testing.T) {
	is := assert.New(t)
	ch := make(chan int)
	is.Equal(IsChannel(ch), true)
}

func TestIsSlice(t *testing.T) {
	is := assert.New(t)
	s := make([]string, 0)
	is.Equal(IsSlice(s), true)
}

func TestIsTime(t *testing.T) {
	is := assert.New(t)
	now := time.Now()
	is.Equal(IsTime(now), true)
}

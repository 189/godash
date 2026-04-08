package godash

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDuration(t *testing.T) {
	is := assert.New(t)

	d := time.Second*3 + time.Hour*2 + time.Minute*21
	str := FormatDuration(d, "HH:mm:ss")

	expected := "02:21:03"
	is.Equal(expected, str)
}

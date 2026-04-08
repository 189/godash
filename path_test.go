package godash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixWindowsPath(t *testing.T) {
	is := assert.New(t)

	data := make(map[string]string, 0)
	data["Name"] = "Tom"
	data["Country"] = "USA"

	rst := FixWindowsPath("/c/a/../a/./c")
	expected := "c://a/c"

	if !is.Equal(rst, expected) {
		t.Fatalf("expect %v, but got %v", expected, rst)
	}
}

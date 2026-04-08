package godash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPadLeft(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ret := PadStart("x", 3, "0")
	expect := "00x"

	is.Equal(expect, ret)
}

func TestPadEnd(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ret := PadEnd("x", 3, "0")
	expect := "x00"

	is.Equal(expect, ret)
}

func TestSubstring(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ret := Substring("hello world", 3, 3)
	expect := "lo "

	is.Equal(expect, ret)
}

func TestCapitalize(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ret := Capitalize("hello world")
	expect := "Hello World"

	is.Equal(expect, ret)
}

func TestPascalCase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ret := PascalCase("hello world")
	expect := "HelloWorld"

	is.Equal(expect, ret)

}

func TestCamelCase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ret := CamelCase("hello world")
	expect := "helloWorld"

	is.Equal(expect, ret)
}

func TestKebabCase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ret := KebabCase("hello world")
	expect := "hello-world"

	is.Equal(expect, ret)
}

func TestIsAlpha(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rst := IsAlpha("13131")
	is.Equal(false, rst)

	rst = IsAlpha("abcdef")
	is.Equal(true, rst)
}

func TestIsAlphaNumberic(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rst := IsAlphaNumeric("abdefg1331$")
	is.Equal(false, rst)

	rst = IsAlphaNumeric("abdefg1331")
	is.Equal(true, rst)

	rst = IsAlphaNumeric("abcd")
	is.Equal(true, rst)

	rst = IsAlphaNumeric("1325")
	is.Equal(true, rst)
}

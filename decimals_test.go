package godash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimalsMult(t *testing.T) {
	is := assert.New(t)

	v, _ := DecimalsMult("0.001", "1000000000000000000")
	is.Equal("1000000000000000", v)

	v, _ = DecimalsMult("0.001", "1000000000000000000", "0.1")
	is.Equal("100000000000000", v)

	v, _ = DecimalsMult("0.001", "1000000000000000000", "0.1", "10")
	is.Equal("1000000000000000", v)
}

func TestDecimalsPlus(t *testing.T) {
	is := assert.New(t)

	v, _ := DecimalsPlus("0.1", "0.2")
	is.Equal("0.3", v)

	v, _ = DecimalsPlus("0.1", "0.2", "0.3")
	is.Equal("0.6", v)

	v, _ = DecimalsPlus("10", "0.2", "31")
	is.Equal("41.2", v)

	_, err := DecimalsPlus("0.1")
	is.Error(err)
}

func TestDecimalsDiv(t *testing.T) {
	is := assert.New(t)

	v, _ := DecimalsDiv("100", "0.2", "5")
	is.Equal("100", v)
}

func TestDecimalsDivRound(t *testing.T) {
	is := assert.New(t)

	v, _ := DecimalsDivRound("100", "0.2", 5)
	is.Equal("500", v)

	v, _ = DecimalsDivRound("0.01", "10", 5)
	is.Equal("0.001", v)

	v, _ = DecimalsDivRound("0.006", "10", 3)
	is.Equal("0.001", v)

	v, _ = DecimalsDivRound("0.006", "10", 2)
	is.Equal("0", v)
}

func TestDecimalsPow(t *testing.T) {
	is := assert.New(t)

	v, _ := DecimalsPow("0.1", "2")
	is.Equal("0.01", v)
}

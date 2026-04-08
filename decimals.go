package godash

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type operationFunc func(decimal.Decimal, decimal.Decimal) decimal.Decimal

func calculate(input string, rest []string, operation operationFunc) (string, error) {
	if len(rest) == 0 {
		return "", fmt.Errorf("more arguments are needed")
	}

	total, err := decimal.NewFromString(input)
	if err != nil {
		return "", err
	}

	for _, v := range rest {
		v2, err := decimal.NewFromString(v)
		if err != nil {
			return "", err
		}
		total = operation(total, v2)
	}

	return total.String(), nil
}

// DecimalsMult return the product of mutiple big numbers
func DecimalsMult(input string, rest ...string) (string, error) {
	return calculate(input, rest, func(curr, next decimal.Decimal) decimal.Decimal {
		return curr.Mul(next)
	})
}

// DecimalsPlus return the value of multiple parameters added sequentially, DecimalsPlus("0.1", "0.2") = 0.3,
func DecimalsPlus(input string, rest ...string) (string, error) {
	return calculate(input, rest, func(curr, next decimal.Decimal) decimal.Decimal {
		return curr.Add(next)
	})
}

// DecimalsSub return the value of multiple parameters subtracted sequentially. DecimalsSub("0.3", "0.2", "0.05") = 0.05,
func DecimalsSub(input string, rest ...string) (string, error) {
	return calculate(input, rest, func(curr, next decimal.Decimal) decimal.Decimal {
		return curr.Sub(next)
	})
}

// DecimalsDiv return the value of all parameters divide sequentially
func DecimalsDiv(input string, rest ...string) (string, error) {
	return calculate(input, rest, func(curr, next decimal.Decimal) decimal.Decimal {
		return curr.Div(next)
	})
}

// DecimalsDivRound rounds the decimal to places decimal places. If places < 0, it will round the integer part to the nearest 10^(-places).
func DecimalsDivRound(input string, divider string, precision int32) (string, error) {
	total, err := decimal.NewFromString(input)
	if err != nil {
		return "", err
	}

	d, err := decimal.NewFromString(divider)
	if err != nil {
		return "", err
	}
	return total.DivRound(d, precision).String(), nil
}

func DecimalsPow(input string, d2 string) (string, error) {
	total, err := decimal.NewFromString(input)
	if err != nil {
		return "", err
	}

	d, err := decimal.NewFromString(d2)
	if err != nil {
		return "", err
	}
	return total.Pow(d).String(), nil
}

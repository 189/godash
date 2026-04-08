package godash

import "fmt"

func ExampleDecimalsMult() {
	v, _ := DecimalsMult("0.001", "10", "1")
	fmt.Println(v)
	// Output:
	// 0.01
}

func ExampleDecimalsPlus() {
	v, _ := DecimalsPlus("0.1", "0.2", "0.3")
	fmt.Println(v)
	// Output:
	// 0.6
}

func ExampleDecimalsSub() {
	v, _ := DecimalsSub("1", "0.2", "0.3")
	fmt.Println(v)
	// Output:
	// 0.5
}

func ExampleDecimalsDiv() {
	v, _ := DecimalsDiv("100", "0.2", "5")
	fmt.Println(v)
	// Output:
	// 100
}
func ExampleDecimalsDivRound() {
	v, _ := DecimalsDivRound("0.006", "10", 3)
	fmt.Println(v)
	// Output:
	// 0.001
}

func ExampleDecimalsPow() {
	v, _ := DecimalsPow("0.1", "2")
	fmt.Println(v)
	// Output:
	// 0.01
}

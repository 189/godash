package godash

import "fmt"

func ExampleIsAlpha() {
	ret := IsAlpha("basdcf")
	fmt.Println(ret)
	// Output
	// true

}

func ExampleIsAlphaNumeric() {
	rst := IsAlphaNumeric("abcd13")
	fmt.Println(rst)
	// Output:
	// true
}

func ExampleIsNumeric() {
	rst := IsNumeric("1234")
	fmt.Println(rst)
	// Output:
	// true
}

func ExamplePadStart() {
	rst := PadStart("x", 5, "0")
	fmt.Println(rst)
	// Output:
	// 0000x
}

func ExamplePadEnd() {
	rst := PadEnd("x", 5, "0")
	fmt.Println(rst)
	// Output:
	// x0000
}

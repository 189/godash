package godash

import (
	"fmt"
)

func ExampleTry() {

	b := Try(func() error {
		panic("ho no, don't panic here")
	})

	fmt.Println(b)
	// Output:
	// false
}

package godash

import (
	"fmt"
	"time"
)

func ExampleFormatDuration() {
	d := time.Second*3 + time.Hour*2 + time.Minute*21
	str := FormatDuration(d, "HH:mm:ss")

	fmt.Println(str)
	// Output:
	// 02:21:03
}

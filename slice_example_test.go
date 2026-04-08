package godash

import (
	"fmt"
)

func ExampleReduce() {
	nums := []int32{1, 2, 3, 4, 5}
	total := Reduce(nums, 0, func(prev int32, now int32) int32 {
		return prev + now
	})
	fmt.Printf("%v", total)
	// Output:
	// 15
}

func ExampleMap() {
	nums := []int32{1, 2, 3, 4, 5}
	rst := Map(nums, func(item int32) int32 {
		return item * 2
	})

	fmt.Printf("%v", rst)
	// Output:
	// [2 4 6 8 10]
}

func ExampleFilter() {
	nums := []int32{1, 2, 3, 4, 5}
	rst := Filter(nums, func(item int32) bool {
		return item > 4
	})

	fmt.Printf("%v", rst)
	// Output:
	// [5]
}

func ExampleFlat() {
	items := [][]int{
		{1, 2, 3},
		{1},
		{4, 5, 6},
	}
	rst := Flat(items)
	fmt.Printf("%v", rst)
	// Output:
	// [1 2 3 1 4 5 6]
}

func ExamplePush() {
	items := []int{
		1, 2, 3,
	}
	Push(&items, 4, 5)

	fmt.Printf("%v", items)
	// Output:
	// [1 2 3 4 5]
}

func ExamplePop() {
	items := []int{
		1, 2, 3,
	}
	p, _ := Pop(&items)

	fmt.Printf("%v, %v", p, items)
	// Output:
	// 3, [1 2]
}

func ExampleShift() {
	items := []int{
		1, 2, 3,
	}
	f, _ := Shift(&items)

	fmt.Printf("%v, %v", f, items)
	// Output:
	// 1, [2 3]
}

func ExampleSplice() {

	items := []int{
		1, 2, 7, 8, 5,
	}
	Splice(&items, 2, 2, []int{3, 4}...)

	expected := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Printf("%v", expected)
	// Output:
	// [1 2 3 4 5]
}

func ExampleUnshift() {
	items := []int{1, 2, 7, 8, 5}
	Unshift(&items, 0)
	fmt.Println(items)
	// Output:
	// [0 1 2 7 8 5]
}

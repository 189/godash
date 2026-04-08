package godash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Parallel()
	nums := []int32{1, 2, 3, 4, 5}
	total := Reduce(nums, 0, func(prev int32, now int32) int32 {
		return prev + now
	})
	expected := 15
	if total != int32(expected) {
		t.Fatalf(`expect %d, but got %d`, expected, total)
	}
}

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	nums := []int32{1, 2, 3, 4, 5}

	ret := Map(nums, func(item int32) int32 {
		return item * 2
	})

	expected := []int32{2, 4, 6, 8, 10}
	is.ElementsMatch(expected, ret)
}

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	nums := []int32{1, 2, 3, 4, 5}
	ret := Filter(nums, func(item int32) bool {
		return item > 4
	})
	expected := []int32{5}

	is.ElementsMatch(expected, ret)
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	nums := []int32{1, 2, 3, 4}

	rst := Some(nums, func(item int32) bool { return item == 2 })
	is.Equal(true, rst)

	rst = Some(nums, func(item int32) bool { return item == 5 })
	is.Equal(false, rst)
}

func TestEvery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	nums := []int32{1, 2, 3, 4}
	rst := Every(nums, func(item int32) bool { return item > 0 })
	is.Equal(true, rst)

	rst = Every(nums, func(item int32) bool { return item > 2 })
	is.Equal(false, rst)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	nums := []int32{1, 2, 3, 4}
	count := 0
	total := len(nums)

	ForEach(nums, func(idx int, item int32) bool {
		count++
		return true
	})
	if total != count {
		t.Errorf("expect to loop %d times, but got %d", total, count)
		return
	}

	count = 0
	expected := 2

	ForEach(nums, func(idx int, item int32) bool {
		count++
		return item != 2
	})
	if count != expected {
		t.Errorf("expect to loop %d times, but got %d", expected, count)
		return
	}
}

func TestFlat(t *testing.T) {
	t.Parallel()

	is := assert.New(t)
	items := [][]int{
		{1, 2, 3},
		{1},
		{4, 5, 6},
	}
	rst := Flat(items)
	expected := []int{
		1, 2, 3, 1, 4, 5, 6,
	}
	is.Equal(expected, rst)
}

func TestPush(t *testing.T) {
	t.Parallel()

	is := assert.New(t)
	items := []int{
		1, 2, 3,
	}
	Push(&items, 4, 5)
	expected := []int{
		1, 2, 3, 4, 5,
	}
	is.Equal(expected, items)
}

func TestPop(t *testing.T) {
	t.Parallel()

	is := assert.New(t)
	items := []int{
		1, 2, 3,
	}
	f, _ := Pop(&items)

	expected := 3
	is.Equal(expected, f)
	is.Equal(items, []int{
		1, 2,
	})
}

func TestShift(t *testing.T) {
	t.Parallel()

	is := assert.New(t)
	items := []int{
		1, 2, 3,
	}
	f, _ := Shift(&items)

	expected := 1
	is.Equal(expected, f)
	is.Equal(items, []int{
		2, 3,
	})
}

func TestSplice(t *testing.T) {
	t.Parallel()

	is := assert.New(t)
	items := []int{
		1, 2, 7, 8, 5,
	}
	if err := Splice(&items, 2, 2, []int{3, 4}...); err != nil {
		t.Error(err)
		return
	}

	expected := []int{
		1, 2, 3, 4, 5,
	}
	is.Equal(expected, items)
}

func TestUnshift(t *testing.T) {
	t.Parallel()

	is := assert.New(t)
	items := []int{
		1, 2, 7, 8, 5,
	}

	Unshift(&items, 0)

	expected := []int{
		0, 1, 2, 7, 8, 5,
	}
	is.Equal(expected, items)
}

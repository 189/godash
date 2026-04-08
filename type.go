package godash

import (
	"reflect"
	"time"
)

func typeIs(o interface{}, t reflect.Kind) bool {
	valueOf := reflect.ValueOf(o)
	return valueOf.Kind() == t
}

// IsChannel will return true when type of the o is channel
func IsChannel(o interface{}) bool {
	return typeIs(o, reflect.Chan)
}

// IsDate will return true when type of the data is time.Time
func IsTime(o interface{}) bool {
	if _, ok := o.(time.Time); ok {
		return true
	}

	return false
}

// IsSlice will retrun true when type of the o is slice
func IsSlice(o interface{}) bool {
	return typeIs(o, reflect.Slice)
}

// IsArray will retrun true when type of the o is array
func IsArray(o interface{}) bool {
	return typeIs(o, reflect.Array)
}

// IsBool will return true when type of the o is boolean
func IsBool(o interface{}) bool {
	return typeIs(o, reflect.Bool)
}

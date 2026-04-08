package godash

// Retrive all the keys from a map
func GetKeys[K comparable, V any](obj map[K]V) []K {
	ret := make([]K, 0, len(obj))
	for k := range obj {
		ret = append(ret, k)
	}
	return ret
}

// Retrive all the values from a map
func GetValues[K comparable, V any](obj map[K]V) []V {
	ret := make([]V, 0, len(obj))
	for _, v := range obj {
		ret = append(ret, v)
	}
	return ret
}

// Get an empty value
func GetZeroValue[T any]() T {
	var zero T
	return zero
}

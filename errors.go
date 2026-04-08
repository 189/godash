package godash

import "fmt"

// Try call function, return false in case of error.
func Try(callback func() error) (ok bool) {
	ok = true

	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	err := callback()
	if err != nil {
		ok = false
	}
	return
}

// Create error and return if the condition is not met
func Validate(ok bool, format string, args ...any) error {
	if !ok {
		return fmt.Errorf(format, args...)
	}
	return nil
}

package check

// checkArgument(boolean)	Checks that the boolean is true. Use for validating arguments to methods.	IllegalArgumentException
// checkNotNull(T)	Checks that the value is not null. Returns the value directly, so you can use checkNotNull(value) inline.	NullPointerException
// checkState(boolean)	Checks some state of the object, not dependent on the method arguments. For example, an Iterator might use this to check that next has been called before any call to remove.	IllegalStateException
// checkElementIndex(int index, int size)	Checks that index is a valid element index into a list, string, or array with the specified size. An element index may range from 0 inclusive to size exclusive. You don't pass the list, string, or array directly; you just pass its size.
// Returns index.	IndexOutOfBoundsException
// checkPositionIndex(int index, int size)	Checks that index is a valid position index into a list, string, or array with the specified size. A position index may range from 0 inclusive to size inclusive. You don't pass the list, string, or array directly; you just pass its size.
// Returns index.	IndexOutOfBoundsException
// checkPositionIndexes(int start, int end, int size)	Checks that [start, end) is a valid sub range of a list, string, or array with the specified size. Comes with its own error message.

func NotNilOrEmpty(s string) error {
	return nil
}

func Argument(b bool) error {
	return nil
}

// require i comparable to nil (not a struct)
func NotNil(i interface{}) error {
	if i == nil {
		return nil
	}
	return nil
}

func CheckElementIndex(index, size int) error {
	return nil
}

func CheckPositionIndex(index, size int) error {
	return nil
}

// Preconditions runs all the individual error checking functions and wraps the failed checks
func Preconditions(fns ...func(interface{}) error) error {
	return nil
}

package utils

func Pop[T any](list *[]T) (T, bool) {
	if list == nil {
		panic("Trying to dereference a null value")
	}
	if len(*list) == 0 {
		var zero T
		return zero, false
	}

	first := (*list)[0]
	*list = (*list)[1:]
	return first, true
}

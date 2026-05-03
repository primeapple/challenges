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

func Filter[T any](list []T, f func(T) bool) []T {
	filtered := make([]T, 0)
	for _, item := range list {
		if f(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func Map[IN, OUT any](list []IN, f func(IN, int) OUT) []OUT {
	var mapped = make([]OUT, len(list))
	for i, item := range list {
		mapped[i] = f(item, i)
	}
	return mapped
}

func Reduce[IN, OUT any](list []IN, accumulator OUT, f func(IN, OUT) OUT) OUT {
	for _, item := range list {
		accumulator = f(item, accumulator)
	}
	return accumulator
}

func Unique[T comparable](list []T) []T {
	seen := map[T]struct{}{}
	result := []T{}
	for _, item := range list {
		_, exists := seen[item]
		if !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

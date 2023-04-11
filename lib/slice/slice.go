package slice

func Reduce[T, U any](slice []T, reducer func(acc U, curr T) U, initialValue U) U {
	result := initialValue
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

func Map[T, U any](slice []T, mapper func(T) U) []U {
	mapped := make([]U, len(slice))
	for i, v := range slice {
		mapped[i] = mapper(v)
	}
	return mapped
}

func IndexOf[T any](slice []T, cmp func(T) bool) int {
	for i, v := range slice {
		if cmp(v) {
			return i
		}
	}

	return -1
}

func Filter[T any](slice []T, cmp func(T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range slice {
		if cmp(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

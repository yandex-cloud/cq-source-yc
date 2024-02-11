package util

func SliceToSet[T comparable](src []T) map[T]bool {
	result := make(map[T]bool)
	for _, elem := range src {
		result[elem] = true
	}
	return result
}

func SetToSlice[T comparable](src map[T]bool) []T {
	result := make([]T, 0)
	for key := range src {
		result = append(result, key)
	}
	return result
}

func Filter[T comparable](arr []T, predicate func(T) bool) []T {
	filtered := make([]T, 0)
	for _, elem := range arr {
		if predicate(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

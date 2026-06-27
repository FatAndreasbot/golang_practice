package t3

func RemoveDuplicates[T comparable](s []T) []T {
	uniqueSet := make(map[T]bool)
	for _, v := range s {
		uniqueSet[v] = true
	}
	unique := make([]T, 0, len(uniqueSet))

	for v := range uniqueSet {
		unique = append(unique, v)
	}

	return unique
}

func RemoveIf[T any](s []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range s {
		if !predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

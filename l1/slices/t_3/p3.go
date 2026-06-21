package t3

func reduceCapacity[T any](s []T) []T {
	if cap(s)/2 < len(s) {
		return s
	}
	reducedSlice := make([]T, 0, cap(s)/2)
	copy(reducedSlice, s)
	return reducedSlice
}

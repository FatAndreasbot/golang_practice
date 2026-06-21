package t3

func RemoveOrdered[T any](s []T, i int) []T {
	s = append(s[0:i], s[i+1:]...)

	return reduceCapacity(s)
}

func RemoveUnordered[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]

	return reduceCapacity(s)
}

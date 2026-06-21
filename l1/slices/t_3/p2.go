package t3

func RemoveAllBtValue[T comparable](s []T, value T) []T {
	s2 := make([]T, 0, len(s))

	for _, v := range s {
		if v != value {
			s2 = append(s2, v)
		}
	}
	return s2

}

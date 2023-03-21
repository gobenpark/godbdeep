package internal

func Index[T comparable](a []T, b T) int {
	for i := range a {
		if a[i] == b {
			return i
		}
	}
	return -1
}

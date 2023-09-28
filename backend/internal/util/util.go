package util

func Ptr[T any](v T) *T {
	return &v
}

func Contains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

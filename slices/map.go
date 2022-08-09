package slices

func Map[T, R any](slice []T, fn func(T) R) []R {
	res := make([]R, len(slice))

	for i, v := range slice {
		res[i] = fn(v)
	}

	return res
}

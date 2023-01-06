package order

func Filter[T any](input []T, f func(T) bool) []T {
	results := make([]T, 0)
	for _, item := range input {
		if f(item) {
			results = append(results, item)
		}
	}
	return results
}

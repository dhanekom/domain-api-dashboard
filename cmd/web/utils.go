package main

func filter[T any](values []T, test func(T) bool) []T {
	var result []T
	for _, value := range values {
		if test(value) {
			result = append(result, value)
		}
	}

	return result
}

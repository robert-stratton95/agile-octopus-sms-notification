package utils

func MapSlice[T any, U any](slice []T, mappingFunc func(T) U) []U {
	resultSlice := make([]U, 0)
	for _, t := range slice {
		resultSlice = append(resultSlice, mappingFunc(t))
	}
	return resultSlice
}

func FilterSlice[T any](list []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, t := range list {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}

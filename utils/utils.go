package utils

/*
Given a predicate and a list, returns if a value passing the
predicate exists
*/
func ExistsInList[T any](pred func(T) bool, list []T) bool {
	for _, value := range list {
		if pred(value) {
			return true
		}
	}

	return false
}

func FilterList[T any](pred func(T) bool, list []T) []T {
	result := []T{}

	for _, value := range list {
		if pred(value) {
			result = append(result, value)
		}
	}

	return result
}

func MapList[T any, U any](transform func(T) U, list []T) []U {
	result := []U{}

	for _, value := range list {
		result = append(result, transform(value))
	}

	return result
}

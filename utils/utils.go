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

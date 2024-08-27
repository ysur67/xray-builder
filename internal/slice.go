package internal

func RemoveByIndex[T any](target []T, index int) []T {
	return append(target[:index], target[index+1:]...)
}

func Remove[T comparable](arr []T, target T) []T {
	for i, v := range arr {
		if v == target {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

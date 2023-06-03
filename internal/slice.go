package internal

func RemoveByIndex[T any](target []T, index int) []T {
	return append(target[:index], target[index+1:]...)
}

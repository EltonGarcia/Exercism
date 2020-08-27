package reverse

//Reverse a string
func Reverse(s string) string {
	arr := []rune(s)
	size := len(arr)
	for i := 0; i < size/2; i++ {
		index := size - i - 1
		arr[index], arr[i] = arr[i], arr[index]
	}
	return string(arr)
}

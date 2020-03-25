package main

import (
	"fmt",
	"strings"
)

func reverseWords(s string) string {

	i := 0
	j := len(s) - 1

	b := []byte(s)

	reverse(b)

	i = 0
	j = 0
	
	fmt.Printf("0| %s\n", b)

	for i < len(b) {
		if b[i] == ' ' {
			i++
			j++
		} else if j < len(b) && b[j] != ' ' {
			j++
		} else {
			reverse(b[i:j])
			i = j
		}
	}

	return strings.TrimSpace(string(b))
}

func reverse(arr []byte) {

	i := 0
	j := len(arr) - 1

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {

	var input string
	input = "I am a student"
	fmt.Printf("input:%s, output:%s\n", input, reverseWords(input))

	input = "am"
	fmt.Printf("input:%s, output:%s\n", input, reverseWords(input))

	input = "am "
	fmt.Printf("input:%s, output:%s\n", input, reverseWords(input))

	input = "  hello world!  "
	fmt.Printf("input:%s, output:%s\n", input, reverseWords(input))

}
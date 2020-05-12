package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	m := make(map[byte]byte)

	var ans, i, j int

	for j < len(s) {

		if _, ok := m[s[j]]; ok {
			delete(m, s[i])
			i++
		} else {

			m[s[j]] = s[j]
			j++

			if j - i > ans {
				ans = j - i
			}
		}
	}

	return ans

}

func main() {

	fmt.Printf("%d\n", lengthOfLongestSubstring("abcabcdbb"))
}
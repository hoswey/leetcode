package main

import "fmt"
// 实现 strStr() 函数。

// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

// 示例 1:

// 输入: haystack = "hello", needle = "ll"
// 输出: 2
// 示例 2:

// 输入: haystack = "aaaaa", needle = "bba"
// 输出: -1
// 说明:
func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := range haystack {
		
		if haystack[i] == needle[0] {

			same := true
			for j := 1; j < len(needle); j++ {

				if i + j >= len(haystack) {
					same = false
					break
				}

				if haystack[i + j] != needle[j] {
					same = false
					break
				}
			}

			if same {
				return i
			}
		}
	}

	return -1
}

func main() {

	var haystack, needle string
	haystack = "hello"
	needle = "ll"

	fmt.Printf("haystack:%s,needle:%s, output:%d, expected:%d\n", haystack, needle, strStr(haystack, needle), 2)

	haystack = "aaaaa"
	needle = "bba"
	fmt.Printf("haystack:%s,needle:%s, output:%d, expected:%d\n", haystack, needle, strStr(haystack, needle), -1)

}


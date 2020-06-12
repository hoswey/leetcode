package main

import "fmt"
// 一条包含字母 A-Z 的消息通过以下方式进行了编码：

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。

// 示例 1:

// 输入: "12"
// 输出: 2
// 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
// 示例 2:

// 输入: "226"
// 输出: 3
// 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。


func numDecodings(s string) int {

	if len(s) == 0 || s[0] == '0'{
		return 0
	}

	dp := make([]int, len(s) + 1)
	dp[0] = 1
	dp[1] = 1

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] >= '1' && s[i] <= '6' {
			dp[i+1] = dp[i] + dp[i-1]
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[len(s)]
}


func main() {

	var s string

	s = "10"
	fmt.Printf("s:%v,output:%d,expect:%d\n", s, numDecodings(s), 1)	

	s = "12"
	fmt.Printf("s:%v,output:%d,expect:%d\n", s, numDecodings(s), 2)

	s = "226"
	fmt.Printf("s:%v,output:%d,expect:%d\n", s, numDecodings(s), 3)	

	s = "0"
	fmt.Printf("s:%v,output:%d,expect:%d\n", s, numDecodings(s), 0)	

	s = "10100"
	fmt.Printf("s:%v,output:%d,expect:%d\n", s, numDecodings(s), 0)		
}
package main

import "fmt"
// 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

// 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

// L   C   I   R
// E T O E S I I G
// E   D   H   N
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

// 请你实现这个将字符串进行指定行数变换的函数：

// string convert(string s, int numRows);

// 示例 1:
// 输入: s = "LEETCODEISHIRING", numRows = 3
// 输出: "LCIRETOESIIGEDHN"


// 示例 2:
// 输入: s = "LEETCODEISHIRING", numRows = 4
// 输出: "LDREOEIIECIHNTSG"
// 解释:

// L     D     R
// E   O E   I I
// E C   I H   N
// T     S     G

func convert(s string, numRows int) string {

	if len(s) == 0 || numRows == 0 {
		return ""
	}

	if numRows == 1 {
		return s
	}

	arr := make([][]byte, numRows)
	for i := range arr {
		arr[i] = make([]byte, len(s))
	}

	chars := []byte(s)

	down := true
	r := -1
	c := 0
	for i := 0; i < len(chars); i++ {
		if down {
			r++
			arr[r][c] = chars[i]
			if r == numRows - 1 {
				down = false
			} 
		} else {
			r--
			c++
			arr[r][c] = chars[i]
			if r == 0 {
				down = true
			}
		}
	}

	var ans []byte
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] != 0 {
				ans = append(ans, arr[i][j])
			}
		}
	}
	return string(ans)
}
 
func main() {

	var input string
	var numRows int

	input = "LEETCODEISHIRING"
	numRows = 3
	fmt.Printf("input:%s, numRows:%d,output:%s\n", input, numRows, convert(input, numRows))

	numRows = 4
	fmt.Printf("input:%s, numRows:%d,output:%s\n", input, numRows, convert(input, numRows))

	numRows = 1
	fmt.Printf("input:%s, numRows:%d,output:%s\n", input, numRows, convert(input, numRows))

	input = ""
	numRows = 1
	fmt.Printf("input:%s, numRows:%d,output:%s\n", input, numRows, convert(input, numRows))
}

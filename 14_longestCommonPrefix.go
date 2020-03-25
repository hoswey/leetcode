package main
import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	l := minLen(strs)
	if l == 0 {
		return ""
	}

	var i int
	for ;i < l; i++ {

		same := true

		for j := 0; j < len(strs) - 1; j++ {
			if strs[j][i] != strs[j+1][i] {
				same = false
				break
			}
		}

		if !same {
			break
		}
	}

	return strs[0][:i]
}

func minLen(strs []string) int {

	min := 1 << 31 - 1
	for i := range strs {

		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}

	return min
}

func main() {

	var strs []string
	strs = []string{"xxower","flow","flight"}
	fmt.Printf("%v\n", longestCommonPrefix(strs))
}


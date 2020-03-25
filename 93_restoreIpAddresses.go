package main

import (
	"strconv"
	"fmt"
)
// 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

// 示例:

// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]

func restoreIpAddresses(s string) []string {
    
    var path []string
    var ans [][]string
	for i := 1; i <= 3; i++ {
		if len(s) >= i {
			backtracking(s[i:], s[0:i], 0, path, &ans)
		}		
	}    

	var fmtAns []string
	for i := range ans {
		var ip string
		for j := range ans[i] {
			ip = ip + ans[i][j]
			if j < len(ans[i]) - 1 {
				ip += "."
			}
		}
		fmtAns = append(fmtAns, ip)
	}

	return fmtAns
}

func backtracking(remain string, part string, size int, path []string, ans *[][]string){

	if !validIp(part) {
		return
	}

	if size == 3 {
		if remain == "" {
			path = append(path, part)
			*ans = append(*ans, path)
		}
	} else {
		for i := 1; i <= 3; i++ {
			if len(remain) >= i {
				newPath := append(path, part)
                newSize := size + 1
				backtracking(remain[i:], remain[0:i], newSize, newPath, ans)
			}		
		}		
	}
}

func validIp(part string) bool{

	if len(part) > 1 && part[0] == '0' {
		return false
	}


	i, err := strconv.Atoi(part)
	if err != nil {
		return false
	}

	if i >= 0 && i <= 255 {
		return true
	}

	return false
}

func main() {

	s := "010010"
	fmt.Printf("%v", restoreIpAddresses(s))
}
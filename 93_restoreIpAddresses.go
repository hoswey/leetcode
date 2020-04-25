package main

import (
	"strconv"
	"fmt"
)
// 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

// 示例:

// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]

import "strconv"

func restoreIpAddresses(s string) []string {

    var ans []string
    backtrack(0, s, nil, &ans)
    return ans
}

func backtrack(pos int, s string, ip []string, ans *[]string) {

    for i := 1; i <= 3; i++ {
        if pos + i > len(s) {
            return
        }

        part := s[pos: pos+i]
        if isValidIp(part) {

            ipCp := make([]string, len(ip))
            copy(ipCp, ip)
            ipCp = append(ipCp, part)

            if len(ipCp) == 4 {
                if pos+i == len(s) {
                    *ans = append(*ans, convertToIpString(ipCp))
                }
            } else {
                backtrack(pos + i, s, ipCp, ans)
            }
        }
    }
}

func convertToIpString(ips []string) string {

    var str string
    for i := 0; i < len(ips); i++ {
        str += ips[i]
        if i != len(ips) - 1 {
            str += "."
        }
    }
    return str
}


func isValidIp(s string) bool {

    if len(s) > 1 && s[0] == '0' {
        return false
    }

    i, err := strconv.Atoi(s)
    if err != nil {
        return false
    }

    return i >= 0 && i <= 255
}

func main() {

	s := "010010"
	fmt.Printf("%v", restoreIpAddresses(s))
}

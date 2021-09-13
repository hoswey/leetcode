/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
import "strings"
import "strconv"

func restoreIpAddresses(s string) []string {

	//Input: s = "25525511135"
	//Output: ["255.255.11.135","255.255.111.35"]

	var ans []string

	var backtracking func(string, []string)
	backtracking = func(s string, parts []string) {

		if len(parts) == 4 {
			if len(s) == 0 {
				ans = append(ans, strings.Join(parts, "."))
			}
			return
		}

		if len(s) >= 1 {
			backtracking(string(s[1:]), append(parts, string(s[:1])))
		}

		if len(s) >= 2 && s[0] != '0' {
			backtracking(string(s[2:]), append(parts, string(s[:2])))
		}

		if len(s) >= 3 && s[0] != '0' {
			num, _ := strconv.Atoi(string(s[0:3]))

		 	if num <= 255 {
				backtracking(string(s[3:]), append(parts, string(s[:3])))
			}
		}
	}
	backtracking(s, nil)
	return ans
}
// @lc code=end


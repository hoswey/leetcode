
// 输入 "1212"
// 输出 4
// 预期结果 5
//变态的题目输入， 可能输入各种0， 对于无法解码的字符串得返回0， 例如连续两个0
func numDecodings(s string) int {
 
	if s[0] == '0' {
		return 0
	}
	
	digits := make([]int, len(s))

	for i, c := range s {	
		digits[i] = int(c - '0')
	}

	for i, _ := range digits {	
		if i < len(s) - 1 {
			if digits[i] == 0 && digits[i+1] == 0 {
				return 0
			}
		}
	}
	
	dp := make([]int, len(s))
	dp[0] = 1

	for i := 1; i < len(digits); i++ {
		 
		validC := (digits[i-1] * 10 + digits[i]) <= 26
		//当前的0无法和上一个0组合
		if digits[i] == 0 && !validC {
			return 0
		}

		if digits[i-1] == 0 {
		  //上一个值为0，所以上一个值固定的得和上上个值合并		
		   dp[i] = dp[i-1]
		} else if digits[i] != 0 && !validC {
		//当前的值不等于0，和上一个数据也无法合并，那么也就是i-1
		   dp[i] = dp[i-1]
		} else if digits[i] != 0 && validC {
		//当前的值不等于0，也可以和上个数字合并，那么就是上两个值的dp和
			if i >= 2 {
				 dp[i] = dp[i-1] + dp[i-2] 
			 } else {
				 dp[i] = dp[i-1] + 1
			 }
		 } else {
		//当前值为0，那么固定的和上个值合并
			if i >= 2 {
				 dp[i] = dp[i-2] 
			 } else {
				 dp[i] = 1
			 }			 
		 }
	 }

	 fmt.Printf("dp:%v\n", dp)

	 return dp[len(digits) - 1]
}
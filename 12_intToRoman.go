/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {

	vals := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	chars := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}

	var ans string
	for num != 0 {
		for i, val := range vals {
			if num >= val {
				ans += chars[i]
				num -= val
				break
			}
		}
	}

	return ans
}
// @lc code=end


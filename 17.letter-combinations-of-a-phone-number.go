/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}

	chars := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var ans []string
	var recurse func(int, []byte)

	recurse = func(from int, bytes []byte) {

		if from == len(digits) {
			ans = append(ans, string(bytes))
			return
		}

		digit := chars[int(digits[from]) - '0' - 2]
		for i := 0; i < len(digit); i++ {
			recurse(from+1, append(append([]byte(nil), bytes...), digit[i]))
		}
	}

	recurse(0, nil)
	return ans
}
// @lc code=end
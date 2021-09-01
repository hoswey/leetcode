/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 */

// @lc code=start
func maxProduct(words []string) int {

	wbits := make([]int, len(words))

	for i, word := range words {
		for j := 0; j < len(word); j++ {
			wbits[i] = wbits[i] | (1 << (word[j] - 'a'))
		}
	}

	var ans int
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if wbits[i]&wbits[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


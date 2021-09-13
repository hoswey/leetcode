/*
 * @lc app=leetcode id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 */

// @lc code=start
func findKthNumber(n int, k int) int {

	prefix := 1
	p := 1

	for p < k {

		count := getCount(prefix, n)

		if p+count > k {
			p++
			prefix *= 10
		} else {
			p += count
			prefix += 1
		}
	}

	return prefix
}

func getCount(prefix, n int) int {

	cur := prefix
	next := cur + 1

	var count int
	for cur <= n {

		count += min(n+1, next) - cur
		cur *= 10
		next *= 10
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func hammingDistance(x int, y int) int {
    diff := x ^ y

	var ans int
	for i := 0; i < 32; i++ {
		if diff & 1 == 1 {
			ans++
		}
		diff = diff >> 1
	}
	return ans
}
// @lc code=end


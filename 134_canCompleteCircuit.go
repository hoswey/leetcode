/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {

	var start, run, rest int

	for i := range gas {

		run += gas[i] - cost[i]
		rest +=  gas[i] - cost[i]

		if run < 0 {
			run = 0
			start = i + 1 
		}
	}

	if rest < 0 {
		return -1
	} else {
		return start
	}
}
// @lc code=end


/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {

		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {

		last := ans[len(ans)-1]

		if last[1] >= intervals[i][0] {

			var right int
			if last[1] > intervals[i][1] {
				right = last[1]
			} else {
				right = intervals[i][1]
			}
			ans[len(ans)-1] = []int{last[0], right}
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

// @lc code=end


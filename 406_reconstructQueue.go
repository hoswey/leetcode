/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
import "sort"

func reconstructQueue(people [][]int) [][]int {
	
	sort.Slice(people, func(i, j int) bool {
		fir, sec := people[i], people[j]
		if fir[0] == sec[0] {
			return fir[1] > sec[1]
		} else {
			return fir[0] < sec[0]
		}
	})

	ans := make([][]int, len(people))
	for _, p := range people {
		k := p[1]
		var space int
		for i := 0; i < len(ans); i++ {
			if ans[i] == nil {
				if space == k  {
					ans[i] = p
					break
				}
				space++
			}
		}
	}
	return ans
}
// @lc code=end


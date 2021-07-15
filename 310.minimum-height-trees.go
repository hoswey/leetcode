/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {

	if n == 1 {
		return []int{0}
	}

	degrees := make([]int, n)
	neighbors := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		l, r := edges[i][0], edges[i][1]
		degrees[l]++
		degrees[r]++
		neighbors[l] = append(neighbors[l], r)
		neighbors[r] = append(neighbors[r], l)
	}

	q := list.New()
	for i, degree := range degrees {
		if degree == 1 {
			q.PushBack(i)
		}
	}

	var ans []int
	for q.Len() > 0 {

		ans = nil
		for i := q.Len(); i > 0; i-- {
			e := q.Front()
			q.Remove(e)

			v := e.Value.(int)
			ans = append(ans, v)
			for j := 0; j < len(neighbors[v]); j++ {
				degrees[neighbors[v][j]]--
				if degrees[neighbors[v][j]] == 1 {
					q.PushBack(neighbors[v][j])
				}
			}
		}
	}
	return ans
}
// @lc code=end


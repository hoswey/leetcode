/*
 * @lc app=leetcode id=547 lang=golang
 *
 * [547] Number of Provinces
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {

	visit := make([]bool, len(isConnected))

	var dfs func(int)
	dfs = func(from int) {
		visit[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !visit[to] {
				dfs(to)
			}
		}
	}

	var ans int
	for i, v := range visit {
		if !v {
			dfs(i)
			ans++
		}
	}
	return ans
}
// @lc code=end


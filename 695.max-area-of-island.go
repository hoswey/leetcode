/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	var ans int
	visit := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visit[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			var cur int
			dfs(grid, visit,i, j, &cur)
			if cur > ans {
				ans = cur
			}
		}
	}
	return ans
}

func dfs(grid [][]int, visit [][]bool, i, j int, cur *int) {

	if i < 0 || i >= len(grid) {
		return
	}

	if j < 0 || j >= len(grid[i]) {
		return
	}

	if visit[i][j] || grid[i][j] == 0 {
		return
	}

	visit[i][j] = true
	(*cur)++

	dfs(grid, visit, i - 1, j, cur)
	dfs(grid, visit, i + 1, j, cur)
	dfs(grid, visit, i, j - 1, cur)
	dfs(grid, visit, i, j + 1, cur)
}
// @lc code=end


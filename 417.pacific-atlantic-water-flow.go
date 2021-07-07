/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {

	if len(heights) == 0 {
		return nil
	}

	copyGrid := func() [][]int {
		
		grid := make([][]int, len(heights))
		for i := 0; i < len(heights); i++ {
			grid[i] = make([]int, len(heights[i]))
		}
		return grid
	}

	visit := copyGrid()
	flags := copyGrid()

	var dfs func([][]int, int, int,[][]int, [][]int)

	dfs = func(heights [][]int, i, j int, visit [][]int, flags [][]int) {

		if visit[i][j] == 1 {
			return
		}

		visit[i][j] = 1
		flags[i][j] += 1
		if i - 1 >= 0 && heights[i - 1][j] >= heights[i][j] {
			dfs(heights, i - 1, j, visit, flags)
		}

		if i + 1 < len(heights) && heights[i + 1][j] >= heights[i][j] {
			dfs(heights, i + 1, j, visit, flags)		
		}

		if j - 1 >= 0 && heights[i][j - 1] >= heights[i][j] {
			dfs(heights, i, j - 1, visit, flags)
		}

		if j + 1 < len(heights[i]) && heights[i][j+1] >= heights[i][j] {
			dfs(heights, i, j + 1, visit, flags)
		}
	}

	R := len(heights)
	C := len(heights[0])

	for j := 0; j < C; j++ {
		dfs(heights, 0, j, visit, flags)
	}

	for i := 1; i < R; i++ {
		dfs(heights, i, 0, visit, flags)
	}

	visit = copyGrid()
	for j := 0; j < C; j++ {
		dfs(heights, R - 1, j, visit, flags)
	}

	for i := 0; i < R - 1; i++ {
		dfs(heights, i, C - 1, visit, flags)
	}

	var ans [][]int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if flags[i][j] >= 2 {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
// @lc code=end


/*
 * @lc app=leetcode id=934 lang=golang
 *
 * [934] Shortest Bridge
 */

// @lc code=start
import "container/list"

func shortestBridge(grid [][]int) int {
	
	if len(grid) == 0 {
		return 0
	}

	M := len(grid)
	N := len(grid[0])

	l := list.New()

	var dfs func([][]int, int, int)
	dfs = func(grid [][]int, i, j int) {
		if i < 0 || i >= M || j < 0 || j >= N {
			return
		}

		if grid[i][j] == 0 || grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		l.PushBack([]int{i, j})
		dfs(grid, i - 1, j)
		dfs(grid, i + 1, j)
		dfs(grid, i, j - 1)
		dfs(grid, i, j + 1)
	}

	var islandFound bool
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {

			if islandFound {
				break
			}

			if grid[i][j] == 1 {
				dfs(grid, i, j)
				islandFound = true
				break
			}
		}
	}

	var level int
	direction := []int{-1, 0, 1, 0, -1}
	for l.Len() > 0 {

		len := l.Len()
		for i := 0; i < len; i++ {

			e := l.Front()
			l.Remove(e)
			cell := e.Value.([]int)

			for k := 0; k < 4; k++ {
				
				r := cell[0] + direction[k]
				c := cell[1] + direction[k+1]
				if r < 0 || r >= M || c < 0 || c >= N {
					continue
				}

				if grid[r][c] == 2 {
					continue
				}

				if grid[r][c] == 1 {
					return level
				}

				if grid[r][c] == 0 {
					l.PushBack([]int{r,c})
				}

				grid[r][c] = 2
			}
		}
		level++
	}

	return 0
}
// @lc code=end


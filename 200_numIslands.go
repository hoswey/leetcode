package main
import "fmt"

// 给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

// 示例 1:

// 输入:
// 11110
// 11010
// 11000
// 00000

// 输出: 1
// 示例 2:

// 输入:
// 11000
// 11000
// 00100
// 00011

// 输出: 3 
func dfs(i, j int, count *int, grid [][]byte, visited [][]bool) {

	if i > len(grid) - 1  || i < 0 {
		return
	}

	if j > len(grid[i]) - 1 || j < 0 {
		return
	}

	if !visited[i][j] && grid[i][j] == '1'{
		visited[i][j] = true
		*count++
	
		dfs(i-1, j, count, grid, visited)
		dfs(i, j+1, count, grid, visited)
		dfs(i+1, j, count, grid, visited)
		dfs(i, j-1, count, grid, visited)
	}

}


func numIslands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}

	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	var num int
	for i := range grid {
		for j := range grid[i] {
			var count int
			dfs(i,j, &count, grid, visited)
			if count > 0 {
				num++
			}
		}
	}

	return num
}

func main() {
//[["1","0","1","1","1"],["1","0","1","0","1"],["1","1","1","0","1"]]
	grid := [][]byte {
		{'1','0','1','1','1'},
		{'1','0','1','0','1'},
		// {'1','1','0','0','1'},
		{'1','1','1','0','1'},				
	}

	fmt.Printf("%v\n", numIslands(grid))
}
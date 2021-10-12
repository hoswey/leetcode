/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {

	r := len(matrix)
	if r == 0 {
		return 0
	}

	c := len(matrix[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {

			if j == 0 || matrix[i][j] == '0' {
				matrix[i][j] = matrix[i][j] - '0'
			} else {
				matrix[i][j] = matrix[i][j-1] + 1

			}
		}
	}

	fmt.Printf("%v\n", matrix)

	var ans int
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {

			minWidth := 1<<31 - 1
			for k := i; k >= 0; k-- {
				minWidth = min(minWidth, int(matrix[k][j]))
				ans = max(ans, minWidth*(i-k+1))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end
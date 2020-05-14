package main

func maximalRectangle(matrix [][]byte) int {

	r := len(matrix)

	if r == 0 {
		return 0
	}

	c := len(matrix[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {

			matrix[i][j] = matrix[i][j] - '0'

			if j == 0 || matrix[i][j] == 0 {
				continue
			}
			matrix[i][j] = matrix[i][j-1] + matrix[i][j]
		}
	}

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

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

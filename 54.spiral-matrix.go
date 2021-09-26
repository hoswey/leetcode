/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return nil
	}

	left_d, right_d := 0, 1
	up_d, down_d := 2, 3
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1

	var ans []int

	direction := right_d
	for {
		switch direction {
		case left_d:
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
			if bottom < top {
				goto END
			}
			direction = up_d

		case right_d:
			for i := left; i <= right; i++ {
				ans = append(ans, matrix[top][i])
			}
			top++
			if top > bottom {
				goto END
			}
			direction = down_d
		case up_d:
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
			if left > right {
				goto END
			}
			direction = right_d

		case down_d:
			for i := top; i <= bottom; i++ {
				ans = append(ans, matrix[i][right])
			}
			right--

			if right < left {
				goto END
			}
			direction = left_d
		}
	}

END:
	return ans
}

// @lc code=end


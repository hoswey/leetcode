// 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

// 示例 1:

// 输入:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// 输出: [1,2,3,6,9,8,7,4,5]
// 示例 2:

// 输入:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
//  

func spiralOrder(matrix [][]int) []int {


	if len(matrix) == 0 {
		return nil
	}

	if len(matrix[0]) == 0 {
		return nil
	}

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	var ans []int
	for {

		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--		

		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
		}
		bottom--

		if bottom < top {
			break
		}

		for i := bottom; i >= top; i--{
			ans = append(ans, matrix[i][left])
		}
		left++

		if left > right {
			break
		}
	}

	return ans
}
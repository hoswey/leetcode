/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := low + (high-low)/2

		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

// @lc code=end


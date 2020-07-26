/*
 * @lc app=leetcode.cn id=1471 lang=golang
 *
 * [1471] 数组中的 k 个最强值
 */

// @lc code=start

import (
	"sort"
)

func getStrongest(arr []int, k int) []int {

	sort.Ints(arr)
	m := arr[(len(arr)-1)/2]
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i] - m) > abs(arr[j] - m) || 
			abs(arr[i] - m) == abs(arr[j] - m) && arr[i] > arr[j]
	})
	return arr[:k]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end


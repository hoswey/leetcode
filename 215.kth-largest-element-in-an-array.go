/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an numsay
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {

	j := len(nums) - k
	start := 0
	end := len(nums) - 1
	for {
		p := partition(nums, start, end)
		if p == j {
			return nums[p]
		}
		if p < j {
			start = p + 1
		} else {
			end = p - 1
		}
	}  
}

func partition(nums []int, start, end int) int {

	//get medium of three
	mid := start + (end - start)/2
	if nums[start] > nums[mid] {
		nums[start], nums[mid] = nums[mid], nums[start]
	}

	if nums[start] > nums[end] {
		nums[start], nums[end] = nums[end], nums[start]
	}

	if nums[mid] < nums[end] {
		nums[mid], nums[end] = nums[end], nums[mid]
	}

	j := start
	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[end] = nums[end],nums[j]
	return j
}
// @lc code=end
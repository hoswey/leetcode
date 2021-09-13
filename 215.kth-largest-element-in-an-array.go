/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an numsay
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {

	//[7,6,5,4,3,2,1]
	//2
	if k > len(nums) {
		return -1
	}

	j := len(nums) - k
	start := 0
	end := len(nums) - 1
	for {

		p := partition(nums, start, end)
		if p == j {
			return nums[p]
		} else if p < j {
			start = p + 1
		} else {
			end = p - 1
		}
	}
}

func partition(nums []int, start, end int) int {

	//medium of three
	mid := start + (end-start)/2
	if nums[start] > nums[mid] {
		nums[start], nums[mid] = nums[mid], nums[start]
	}

	if nums[start] > nums[end] {
		nums[start], nums[end] = nums[end], nums[start]
	}

	if nums[end] > nums[mid] {
		nums[end], nums[mid] = nums[mid], nums[end]
	}

	j := start
	for i := start; i < end; i++ {
		if nums[i] <= nums[end] {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	nums[j], nums[end] = nums[end], nums[j]
	return j
}

// @lc code=end

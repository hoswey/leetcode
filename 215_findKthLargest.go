package main

import "fmt"

// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 示例 1:

// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:

// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
// 说明:

// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度
func findKthLargest(nums []int, k int) int {

	if k > len(nums) {
		return -1
	}

	start := 0
	end := len(nums) - 1

	trg := len(nums) - k
	for {

		p := partition(nums, start, end)

		// start:4, end:8, nums:[2 2 1 3 5 4 6 5 3], p:3, trg:5
		// fmt.Printf("start:%d, end:%d, nums:%v, p:%d, trg:%d\n", start, end, nums,  p, trg)

		if p == trg {
			return nums[p]
		} else if p < trg {
			start = p + 1
		} else {
			end = p - 1
		}
	}
}

func partition(nums []int, start, end int) int {

	mid := start + (end - start)/2

	if nums[start] > nums[mid] {
		nums[start], nums[mid] = nums[mid], nums[start]
	}

	if nums[end] > nums[mid] {
		nums[end], nums[mid] = nums[mid], nums[end]
	} else if nums[mid] < nums[start] {
		nums[mid], nums[start] = nums[start], nums[mid]
	}

	i := start
	for j := start; j < len(nums) - 1; j++ {
		if nums[j] < nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]

	return i
}

func main() {

	var nums []int
	var k int 
	
	nums = []int{3,2,1,5,6,4}
	k = 2

	fmt.Printf("nums:%v, k:%d, output:%d, expected:%d\n", nums, k, findKthLargest(nums,k), 5)

	nums = []int{3,2,3,1,2,4,5,5,6}
	k = 4

	fmt.Printf("nums:%v, k:%d, output:%d, expected:%d\n", nums, k, findKthLargest(nums,k), 4)	


}
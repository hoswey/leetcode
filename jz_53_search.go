// 统计一个数字在排序数组中出现的次数。

//  

// 示例 1:

// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: 2
// 示例 2:

// 输入: nums = [5,7,7,8,8,10], target = 6
// 输出: 0
//  

// 限制：

// 0 <= 数组长度 <= 50000
package main

import "fmt"

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}

	//1 找出左边最后一个小于 target 的数的索引, 有可能找到的数是target
	lo := 0
	hi := len(nums) - 1

	for lo < hi {

		mid := lo + (hi - lo + 1) >> 1
		if nums[mid] < target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	start := lo
	if nums[lo] != target {
		start++
	}

	//2 找出右边第一个大于   target 的数的索引
	lo = 0
	hi = len(nums) - 1 
	for lo < hi {
		mid := lo + (hi - lo) >> 1
		if nums[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	} 

	end := lo
	if nums[lo] != target {
		end--
	}

	l := end - start + 1 
	if l <= 0 {
		return 0
	} 

	return l
}

func main() {

	arr := []int {5, 7,7,8,8,8,10}
	input := 8
	fmt.Printf("input: %d, ouput: %d", input, search(arr, input))
}




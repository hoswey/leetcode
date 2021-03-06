package main

import "fmt"

// 给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

// 示例 1:

// 输入: [1,2,0]
// 输出: 3

// 示例 2:

// 输入: [3,4,-1,1]
// 输出: 2

// 示例 3:
// 输入: [7,8,9,11,12]
// 输出: 1
func firstMissingPositive(nums []int) int {

	if len(nums) == 0 {
		return 1
	}

	for i := 0; i < len(nums); i++ {

		if nums[i] <= 0 || nums[i] > len(nums) {
			continue
		}

		trgIdx := nums[i] - 1

		for i != trgIdx && !(nums[i] <= 0 || nums[i] > len(nums)) {

			if nums[trgIdx] == nums[i] {
				nums[i] = -1
				break
			}

			nums[i], nums[trgIdx] = nums[trgIdx], nums[i]
			trgIdx = nums[i] - 1
		}
	}


	for i := 0; i < len(nums); i++ {

		if nums[i] != i + 1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func main() {

	arr := []int{1,2,0}

	arr = []int{1,2,3,3,3}
	fmt.Printf("input: %v, output:%d\n", arr, firstMissingPositive(arr))
}


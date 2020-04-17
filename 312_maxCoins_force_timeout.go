package main

import "fmt"
import "math/rand"
// 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

// 现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。

// 求所能获得硬币的最大数量。

// 说明:

// 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
// 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
// 示例:

// 输入: [3,1,5,8]
// 输出: 167 
// 解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//      coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

func maxCoins(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	var ans int
	backtrack(nums, 0, 0, &ans)

	return ans
}

func backtrack(nums []int, curentLevel int, currentCoins int, ans *int) {

	if curentLevel == len(nums) - 2 {
		if currentCoins > *ans {
			*ans = currentCoins
		}
		return
	}

	for i := 1; i <= len(nums) - 2; i++ {

		if nums[i] == -1 {
			continue
		}

		left := i - 1
		for nums[left] == -1 {
			left--
		}
		leftNum := nums[left]

		right := i + 1
		for nums[right] == -1 {
			right++
		}
		rightNum := nums[right]

		temp := nums[i]
		nums[i] = -1

		backtrack(nums, curentLevel + 1, currentCoins + leftNum * rightNum, ans)
		nums[i] = temp
	}
}

func main() {

	var input []int
	input = []int{3,1,5,8}

	fmt.Printf("input:%v, output:%v, expect: %d\n", input, maxCoins(input), 167)

	input = []int{3,1}
	fmt.Printf("input:%v, output:%v, expect: %d\n", input, maxCoins(input), 4)

	input = RandArray(11)

	fmt.Printf("input:%v, output:%v, expect: %d\n", input, maxCoins(input), -1)	

}

func RandArray(n int) []int {
	// needed a seed input else it will generate the same number
	rand.Seed(0)

	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

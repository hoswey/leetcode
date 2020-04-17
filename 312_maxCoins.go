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
		return nums[0]
	}

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, len(nums))
	}

	return getMaxCoins(0, len(nums) - 1, nums, memo)
}

func getMaxCoins(i, j int, nums []int, memo [][]int) int{

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	max := 0
	for k := i + 1; k < j; k++ {

		coins := nums[i] * nums[j] * nums[k] + getMaxCoins(i, k, nums, memo) + getMaxCoins(k, j, nums, memo)
		if coins > max {
			max = coins
		}
	}

	memo[i][j] = max
	// fmt.Printf("i:%d, j:%d, coins:%d\n",i, j, max)

	return max
}


func main() {

	var input []int
	// input = []int{3,1,5,8}

	// fmt.Printf("input:%v, output:%v, expect: %d\n", input, maxCoins(input), 167)

	// input = []int{9}
	// fmt.Printf("input:%v, output:%v, expect: %d\n", input, maxCoins(input), 9)

	input =[]int{8,2,6,8,9,8,1,4,1,5,3,0,7,7,0,4,2,2}
	fmt.Printf("input:%v, output:%v, expect: %d\n", input, maxCoins(input), 3446)

	// input = []int{3,1}
	// fmt.Printf("input:%v, output:%v, expect: %d\n", input, maxCoins(input), 4)

	// input = RandArray(10)

	// fmt.Printf("Begin\n")
	// fmt.Printf("input:%v, output:%v, expect: %d\n", input, maxCoins(input), -1)	

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

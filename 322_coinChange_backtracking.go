package main

import "fmt"
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

// 示例 1:

// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3 
// 解释: 11 = 5 + 5 + 1

// 示例 2:
// 输入: coins = [2], amount = 3
// 输出: -1
// 说明:
func coinChange(coins []int, amount int) int {

	if amount == 0 || len(coins) == 0 {
		return -1
	}

	ans := -1
	backtracking(coins, amount, 0, &ans)
	return ans    
}

func backtracking(coins []int, amount int, count int, ans *int) {

	if amount == 0 {
		if *ans == -1 {
			*ans = count
		} else if *ans > count {
			*ans = count
		}
	}

	for _, coin := range coins {
		if amount >= coin {
			backtracking(coins, amount - coin, count + 1, ans)
		}
	}
}

func main() {

	fmt.Printf("%d\n", coinChange([]int{12,2,10}, 11))
}
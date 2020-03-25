package main

import (
	"fmt"
)
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

// 示例 1:

// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3 
// 解释: 11 = 5 + 5 + 1

// 示例 2:
// 输入: coins = [2], amount = 3
// 输出: -1
// 说明:

var MaxInt = 1 << 31 - 1
func coinChange(coins []int, amount int) int {

	if amount == 0 || len(coins) == 0 {
		return 0
	}

	memo := make([]int, amount + 1)

	backtracking(coins, amount, memo)

	return memo[amount]    
}

func backtracking(coins []int, amount int, memo []int) int{

	if memo[amount] != 0 {
		return memo[amount]
	}

	if amount == 0 {
		return 0
	}

	min := MaxInt
	for _, coin := range coins {
		if amount >= coin {
			count := backtracking(coins, amount - coin, memo)
			// fmt.Printf("amount: %d, coin: %d, count: %d\n", amount, coin, count)
			if count != -1 && count + 1 < min {
				min = count + 1
			}
		}
	}

	if min != MaxInt {
		memo[amount] = min
		return min
	} else {
        memo[amount] = -1
		return -1
	}
}

func main() {
	fmt.Printf("ans %d\n", coinChange([]int{1,2,5}, 11))
}
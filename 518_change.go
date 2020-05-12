package main

import "fmt"

// 给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。 
// 示例 1:

// 输入: amount = 5, coins = [1, 2, 5]
// 输出: 4
// 解释: 有四种方式可以凑成总金额:
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
// 示例 2:

// 输入: amount = 3, coins = [2]
// 输出: 0
// 解释: 只用面额2的硬币不能凑成总金额3。
// 示例 3:

// 输入: amount = 10, coins = [10] 
// 输出: 1
func change(amount int, coins []int) int {

	if amount == 0 {
		return 1
	}

	if len(coins) == 0 {
		return 0
	}

	dp := make([][]int, amount + 1)
	for i := range dp {
		dp[i] = make([]int, len(coins) + 1)
	}

	for k := 1; k <= len(coins); k++ {
		dp[0][k] = 1
	}

	for i := 1; i <= amount; i++ {
		for k := 1; k <= len(coins); k++ {
			if i >= coins[k-1] {
				dp[i][k] = dp[i][k-1] + dp[i-coins[k-1]][k]
			} else {
				dp[i][k] = dp[i][k-1]
			}

		}
	}
	
	return dp[amount][len(coins)]
}

func main(){

	var amount int
	var coins []int

	amount = 5
	coins = []int{1, 2, 5}

	fmt.Printf("amount:%d, coins:%v, ouput:%d, expected:%d\n", amount, coins, change(amount, coins), 4)


	amount = 3
	coins = []int{2}
	fmt.Printf("amount:%d, coins:%v, ouput:%d, expected:%d\n", amount, coins, change(amount, coins), 0)


	amount = 10
	coins = []int{10}
	fmt.Printf("amount:%d, coins:%v, ouput:%d, expected:%d\n", amount, coins, change(amount, coins), 1)
}
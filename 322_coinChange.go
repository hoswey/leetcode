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

func coinChange(coins []int, amount int) int {

    MAX_INT := 1 << 31 - 1
    dp := make([]int, amount + 1)
    for i := 1; i <= amount; i++ {

        m := MAX_INT
        for j := 0; j < len(coins); j++ {
            if i >= coins[j] {
                m = min(m, dp[i - coins[j]] + 1)
            }
        }
        dp[i] = m
        fmt.Printf("i:%v,v:%v\n", i, dp[i])
    }

    fmt.Printf("%v\n",dp)
    if dp[amount] == MAX_INT {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	fmt.Printf("ans %d\n", coinChange([]int{1,2,5}, 11))
}

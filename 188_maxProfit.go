func maxProfit(k int, prices []int) int {

    if k >= len(prices) >> 1 {
        sum := 0
        for i := 1; i < len(prices); i++ {
            profit := prices[i] - prices[i-1]
            if profit > 0 {
                sum += profit
            }
        }
        return sum
    }

    tk0 := make([]int, k+1)
    tk1 := make([]int, k+1)

    for i := range tk1 {
        tk1[i] = -1 << 31
    }

    for _, price := range prices {
        for i := 1; i <= k; i++ {
            tk0[i] = max(tk0[i], tk1[i] + price)
            tk1[i] = max(tk1[i], tk0[i-1] - price)
        }
    }
    return tk0[k]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

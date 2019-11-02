func max(a, b int) int {

    if a > b {
        return a
    } else {
        return b
    }
}

func maxCoins(nums []int) int {
    
    A := append([]int{1}, nums...)
    A  = append(A, 1)

    size := len(nums)
    dp := make([]int, size + 2)

    for i := range dp {
        dp[i] = make([]int, size + 2)
    }

    for length = 1; length <= size; l++ {

        for left := 1;  left <= size - length + 1; left++ {

            right := left + length - 1

            for last := left; last <= right; last++ {
                dp[left][right] = max(dp[left][right],  A[left-1] * A[last] * A[right + 1] + dp[left][last-1] + dp[last+1][right])
            }
        }
    }
    return dp[1][size]
}
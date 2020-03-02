func cuttingRope(n int) int {
    
    if n == 2 {
        return 1
    }
    
    if n == 3 {
        return 2
    }
    
    dp := make([]int, n+1)    
    dp[1] = 1
    dp[2] = 2
    dp[3] = 3
    
    for i := 4; i <= n; i++ {             
        var max int
        for j := 1; j <= i/2; j++ {             
            p := dp[j] * dp[i-j]
            if p > max {
                max = p
            }
        }    
        dp[i] = max
    }
    
    return dp[n]
}
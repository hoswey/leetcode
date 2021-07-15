func countBits(num int) []int {
    
    dp := make([]int, len(num) + 1)
    dp[0] = 0

    for i := 1; i <= len(num); i++ {
    	if i & 1 == 0 {
    		dp[i] = dp[i >> 1]
    	} else {
    		dp[i] = dp[i - 1] + 1
    	}
    }
    return dp
}
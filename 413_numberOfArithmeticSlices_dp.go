func numberOfArithmeticSlices(A []int) int {
	
	dp := make([]int, len(A))
	var ans int
	for i := 2; i < len(A); i++ {
		if A[i] - A[i-1] == A[i-1] - A[i-2] {
			dp[i] = 1 + dp[i-1]
			ans += dp[i]
		}
	}
	return ans
}
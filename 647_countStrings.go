func countSubstrings(s string) int {

	n := len(s)
	if n == 0 {
		return 0
	} 

	if n == 1 {
		return 1
	}

	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	ans := n
	for i := 0; i < n - 1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans += 1
		}
	}

	for l := 2; l < n; l++ {

		for i := 0; i < n - l; i++ {

			j := i + l
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				ans += 1
			}
		}
	}

	return ans
}
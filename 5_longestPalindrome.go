func longestPalindrome(s string) string {

	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	var from, to int
	for i := 0; i < n; i++ {
		dp[i][i] = 1
		from = i
		to = i
	}

	for i := 0; i < n - 1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
			max = 2

			from = i
			to = i + 1
		}
	}

	for l := 3; l <= n; l++ {
		for i := 0; i <= n - l; i++ {
			j := i + l - 1
			if s[i] == s[j] && dp[i+1][j-1] != 0 {
				dp[i][j] = dp[i + 1][j - 1] + 2
				if dp[i][j] > from - to + 1 {
					from = i
					to = j
				}
			}
		}
	}
    return s[from:to+1]
}
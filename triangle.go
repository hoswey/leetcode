func minimumTotal(triangle [][]int) int {

	if len(triangle) == 0 {
		return 0
	}

	dp := make([]int, len(triangle))

	var ans int
	for i, row := range triangle {

		m := 1 << 31
		for j := len(row) - 1; j >= 0; j-- {

			num := row[j]
			if j == 0 {
				dp[j] = dp[0] + num
			} else if j == i {
				dp[j] = dp[j-1] + num
			} else {
				dp[j] = min(dp[j-1], dp[j]) + num
			}
			m = min(dp[j], m)
		}
		fmt.Printf("i:%v, dp:%v\n", i, dp)
		ans = m
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}


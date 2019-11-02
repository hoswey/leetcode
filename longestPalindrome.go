func longestPalindrome(s string) string {

	size := len(s)

	if size <= 1 {
		return s
	}

	if size == 2 {

		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}

	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	
	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}

	var from, to int
	for i := 0; i < size - 1; i++ {
		if s[i] == s[i+1] {
			
			dp[i][i+1] = 2

			from = i
			to = i + 1
		}
	}

	for n := 3; n <= size; n++ {
        
		for i := 0; i <= size - n; i++ {

			j := i + n - 1

			if s[i] == s[j] && dp[i+1][j-1] != 0 {
                
				dp[i][j] = dp[i+1][j-1] + 2
                
				if j - i > to - from {
					from = i 
					to = j
				}
			}
		}
	} 
	return s[from: to+1]
}
//")()())"
// 2
// expected 4
func longestValidParentheses(s string) int {
    
    if len(s) <= 1 {
        return 0
    }

    dp := make([][]int, len(s))
    for i, _ := range dp {
    	dp[i] = make([]int, len(s))
    }

    var lvp int
    for i := 0; i < len(s) - 1; i++ {
		if s[i] == '(' && s[i+1] == ')' {
			dp[i][i+1] = 2
			fmt.Printf("i:%d,j:%d\n", i, i+1)
			lvp = 2
		}    	
    }

    if len(s) < 4 {
    	return lvp
    }

    for l := 4; l <= len(s); l = l + 2 {

    	for i := 0; i < len(s) - l + 1; i++ {

    		j := i + l - 1

    		if dp[i][j-2] > 0 && s[j-1] == '(' && s[j] == ')'{
    			dp[i][j] = dp[i][j-2] + 2
    		} else if dp[i+1][j-1] > 0 && s[i] == '(' && s[j] == ')' {
    			dp[i][j] = dp[i+1][j-1] + 2
    		} else {
    			dp[i][j] = 0
    		}
    		lvp = max(lvp, dp[i][j])
    	}
    } 

    for _, arr := range dp {
    	fmt.Printf("%v\n", arr)
    }
    return lvp
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}	
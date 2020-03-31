import "strconv"

func translateNum(num int) int {
    
    if num <= 9 {
        return 1
    }
    
    numstr := strconv.Itoa(num)
    dp := make([]int, len(numstr))
    dp[0] = 1
    
    if isValid(numstr[0], numstr[1]) {
        dp[1] = 2
    } else {
        dp[1] = 1
    }
    
    
    for i := 2; i < len(numstr); i++ {
        if isValid(numstr[i-1], numstr[i]) {
            dp[i] = dp[i-2] + dp[i-1]
        } else {
            dp[i] = dp[i-1]
        }
    }
    
    return dp[len(numstr) - 1]
}

func isValid(a, b byte) bool {
    
    num, err := strconv.Atoi(string([]byte{a,b}))
    if err != nil {
        return false
    }
    
    if num >= 10 && num <= 25 {
        return true
    }
    
    return false
}
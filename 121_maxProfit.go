func maxProfit(prices []int) int {
    
    if len(prices) <= 1 {
        return 0
    }
    
    var ans int
    buy := prices[0]
    
    for i := 1; i < len(prices); i++ {
        
        earn := prices[i] - buy 
        if earn < 0 {
            buy = prices[i]
        } else {
            ans = max(ans, earn)
        }
    }
    
    return ans
}

func max(a, b  int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
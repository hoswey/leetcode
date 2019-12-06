//[1,2,3,0,2]

func maxProfit(prices []int) int {

    prices = append([]int{0}, prices...)

    dpHold := make([]int, len(prices))
    dpCd   := make([]int, len(prices))
    dpFree := make([]int, len(prices))        
    
    dpHold[0] = -(1 << 31 - 1)
    dpCd[0]   = 0
    dpFree[0] = 0

    for i := 1; i < len(prices); i++ {
    	dpHold[i] = max(dpHold[i-1], dpFree[i-1] - prices[i])
    	dpCd[i] = dpHold[i-1] + prices[i]
    	dpFree[i] = max(dpFree[i-1], dpCd[i-1])
    }

    return max(dpFree[len(prices) - 1], dpCd[len(prices) - 1])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
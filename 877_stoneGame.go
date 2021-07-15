type Pair struct {
	F int
	S int
}


func stoneGame(piles []int) bool {
    
    dp := make([][]Pair, len(piles))

    for i := range dp {
    	dp[i] = make([]Pair, len(piles))
    }

    for i := 0; i < len(piles); i++ {
    	dp[i][i] = Pair{F:piles[i]}
    }

    for l := 1; l <= len(piles) - 1; l++ {

    	for i:= 0; i < len(piles) - l; i++ {

    		j := i + l 

    		left := piles[i] + dp[i+1][j].S 
    		right := piles[j] + dp[i][j-1].S 

    		var p Pair
    		if left >= right {
  				p.F = left
  				p.S = dp[i+1][j].F
    		} else {
    			p.F = right
    			p.S = dp[i][j-1].F
    		}
    		dp[i][j] = p
    	}
    }

    for i := range dp {
    	fmt.Printf("%v\n", dp[i])
    }

    return dp[0][len(piles)- 1].F > dp[0][len(piles)- 1].S
}
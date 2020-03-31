func largestRectangleArea(heights []int) int {
    
    var ans int 
    
    for i := 0 ; i < len(heights); i++ {
        
        minHeight := 1 << 31 - 1
        for j := i; j < len(heights); j++ {           
            minHeight = min(minHeight, heights[j])            
            ans = max(ans, minHeight * (j-i+1))
        }
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

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
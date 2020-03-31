func findNthDigit(n int) int {
    
    if n <= 9 {
        return n
    }
    
    n -= 9    
    i := 2
    
    for {     
        
        rangeNum := (powOf10(i) - powOf10(i - 1)) * i        
        if n > rangeNum {
            i++
            n -= rangeNum
            continue
        }
        
        relative := (n - i + 1) / i
        num := powOf10(i - 1) + relative
        index := n % i
        
        if index != 0 {
            index = i - index
        }
        
        var ans int
        for i := 0; i <= index; i++ {
            ans = num % 10
            num = num / 10
        }
        
        return ans                           
    }
}

func powOf10(n int) int {
    
    ret := 1    
    for i := 0; i < n; i++ {
        ret *= 10
    }
    return ret
}
func add(a int, b int) int {
    
    var ret int
    for {
        
        ret    = a ^ b
        carry := a &b
        
        if carry == 0 {
            break
        }
        
        a = ret
        b = carry << 1
    }
    
    return ret
}
func minArray(numbers []int) int {
    
    lo := 0
    hi := len(numbers) - 1
    
    for lo < hi {
        
        mid := lo + (hi - lo)/2
        
        if numbers[mid] > numbers[hi]{
            lo = mid + 1
        } else if numbers[mid] == numbers[hi] {
            hi = hi - 1
        } else {
            hi = mid
        }
    }
    
    return numbers[lo]
}
func missingNumber(nums []int) int {
    
    if len(nums) == 0 {
        return -1
    }
    
    lo := 0 
    hi := len(nums) - 1
    
    for lo <= hi {
        
        mid := lo + (hi-lo) >> 1        
        if nums[mid] == mid {
            lo = mid + 1
        } else if nums[mid] > mid {
            if mid == 0 ||  nums[mid-1] == mid-1 {
                return mid
            } else {
                hi = mid - 1
            }
        }
    }
    
    return lo
}
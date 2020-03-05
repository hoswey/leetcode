func exchange(nums []int) []int {

    if len(nums) <= 1 {
        return nums
    }
    
    iOdd := 0
    iEven := len(nums) - 1
    
    for {
        
        for iOdd < len(nums) && nums[iOdd] & 1 == 1 {
            iOdd++
        }
        
        for iEven >= 0 && nums[iEven] & 1 == 0 {
            iEven--
        }
        
        if iOdd >= iEven {
            break
        }        
        nums[iOdd], nums[iEven] = nums[iEven], nums[iOdd]        
    }
    return nums
}
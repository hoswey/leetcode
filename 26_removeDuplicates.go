func removeDuplicates(nums []int) int {
    
    if len(nums) == 0 {
    	return 0
    }

    l := 1

    cur := nums[0]
    for i := 1; i < len(nums); i++ {
    	if nums[i] != cur {
    		nums[l] = nums[i]
    		cur = nums[i]
    		l++
    	}
    }
    return l
}
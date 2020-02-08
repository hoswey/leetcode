func threeSum(nums []int) [][]int {
    
    if len(nums) < 3 {
        return nil
    }
    
    sort.Ints(nums)
    
    if nums[0] > 0 {
        return nil
    }
    
    var ans [][]int
    for i := 0; i < len(nums) - 2; i++ {
        
        j := i + 1
        k := len(nums) - 1
        
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        for j < k {  
            
            sum := nums[i] + nums[j] + nums[k] 
            if sum == 0 {
                ans = append(ans, []int{nums[i], nums[j], nums[k]})   
                
                j++
                for j < k && nums[j] == nums[j-1] {
                    j++
                }
                
                k--
                for k > j && nums[k] == nums[k+1] {
                    k--
                }
                
                continue
            } else if sum < 0{
                j++
            } else {
                k--
            }            
        }        
    }
    
    return ans
}
func combinationSum(candidates []int, target int) [][]int {
    
    var ans [][]int
    recurse(candidates, target, len(candidates) - 1, &ans, []int{})
    return ans
    
}

func recurse(candidates []int, target int, k int, ans *[][]int, nums []int) {
    
    if target == 0 {
        *ans = append(*ans, nums)        
        return
    }
    //不选
    if k > 0 {
        recurse(candidates, target, k - 1, ans, nums)
    }
    
    //选
    if k >= 0 && target >= candidates[k] {
        
        dst := make([]int, len(nums))
        copy(dst, nums)
        dst = append(dst, candidates[k])
        recurse(candidates, target - candidates[k], k, ans, dst)
    }
}
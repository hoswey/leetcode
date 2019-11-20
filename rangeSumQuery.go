type NumArray struct {

	dp [][]int
    
}


func Constructor(nums []int) NumArray {

	dp := make([][]int, len(nums))
	for i, _ := range dp {
		dp[i] = make([]int, len(nums))
	}

	for i, n := range nums {
		dp[i][i] = n
	}

	for l := 2; k <= len(nums); l++{
		for i := 0; i <= len(nums) - l; i++ {
			dp[i][i+l-1] = dp[i][i+l-2] + dp[i+l-1]
		}
	}

	numArray := NumArray{dp: dp}
	return NumArray
}


func (this *NumArray) SumRange(i int, j int) int {

	if i > j {
		return 0
	}

	return this.dp[i, j]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
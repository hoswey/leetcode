func numberOfArithmeticSlices(A []int) int {
	
	var ans, pre int
	for i := 2; i < len(A); i++ {
		if A[i] - A[i-1] == A[i-1] - A[i-2] {
			cur := 1 + pre
			pre = cur
			ans += cur
		} else {
			pre = 0
		}
	}
	return ans
}
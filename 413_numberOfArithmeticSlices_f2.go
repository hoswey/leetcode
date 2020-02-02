func numberOfArithmeticSlices(A []int) int {

	if len(A) < 3 {
		return 0
	}

	var ans int
	for i := 0; i < len(A) - 2; i++ {

		diff := A[i+1] - A[i]

		for s := i + 2; s < len(A); s++ {
			if s[i] - s[i-1] == diff {
				ans++
			} else {
				break
			}
		}
	}

	return ans
}
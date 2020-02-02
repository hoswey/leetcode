func numberOfArithmeticSlices(A []int) int {

	if len(A) < 3 {
		return 0
	}

	var ans int
	for i := 0; i < len(A) - 2; i++ {

		diff := A[i+1] - A[i]

		for s := i + 2; s < len(A); s++ {

			for j := i + 2; j <= s; j++ {
				if A[j] - A[j-1] != diff {
					break
				}
				if (j == s) {
					ans++
				}
			}
		}
	}

	return ans
}
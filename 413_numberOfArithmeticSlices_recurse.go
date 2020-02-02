func numberOfArithmeticSlices(A []int) int {

	var sum int 

	recurse(&sum, A, len(A) - 1)

	return sum
}

func recurse(sum *int, A []int, i int) int{

	if i < 2 {
		return 0
	}

	if A[i] - A[i-1] == A[i-1] - A[i-2] {
		ap := 1 + recurse(sum, A, i - 1)
		*sum = *sum + ap
		return ap
	} else {
		recurse(sum, A, i -1)
		return 0
	}
}
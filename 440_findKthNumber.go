func findKthNumber(n int, k int) int {

	prefix := 1
	p := 1

	for p < k {

		count := getCount(prefix, n)
		if p + count > k {
			p++
			prefix *= 10
		} else {
			p += count
			prefix++
		}

	}

	return prefix
    
}

func getCount(i, n int) int {

	cur := i
	next := i+1

	var count int
	for cur <= n {
		count += min(next, n+1) - cur
		cur *= 10
		next *= 10
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func findDuplicate(nums []int) int {

	start := 1
	end := len(nums) - 1

	for start < end {

		mid := start + (end - start) / 2

		count := countRange(nums, start, mid)
		if count > (mid - start + 1) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func countRange(nums []int, start, end int) int {

	var count int
	for _, num := range nums {
		if num >= start && num <= end {
			count++
		}
	}
	return count
}
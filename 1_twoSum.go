func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, n := range nums {

		if v, ok := m[n]; ok {
			return []int{v, i}
		} else {
			m[target - n] = i
		}
	}

	panic("no answer")
}
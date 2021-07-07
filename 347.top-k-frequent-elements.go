/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	var arr []int
	for _, num := range nums {
		
		if _, ok := m[num]; !ok {
			m[num] = 0
			arr = append(arr, num)
		}
		m[num] = m[num] + 1
	}    

	// fmt.Printf("1 %v\n", arr)
	sort.Slice(arr, func(i, j int) bool{
		return m[arr[i]] > m[arr[j]]
	})
	// fmt.Printf("2 %v\n", arr)


	return arr[0: k]
}
// @lc code=end


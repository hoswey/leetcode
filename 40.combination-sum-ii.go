/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {

	if len(candidates) == 0 {
		return nil
	}
	uniqNums, counter := buildCounter(candidates)
	fmt.Printf("uniqNums: %v, counter: %v\n", uniqNums, counter)

	var ans [][]int
	backtracking(uniqNums, counter, 0, target, nil, &ans)
	return ans
}

func buildCounter(candidates []int)([]int, []int) {

	sort.Ints(candidates)

	var uniqNums []int
	var counter []int
	
	uniqNums = append(uniqNums, candidates[0])
	counter = append(counter, 1)
	for i := 1; i < len(candidates); i++ {

		last := len(uniqNums) - 1
		if candidates[i] == uniqNums[last] {
			counter[last]++
		} else {
			uniqNums = append(uniqNums, candidates[i])
			counter = append(counter,1)
		}
	}
	return uniqNums, counter
}

func backtracking(uniqNums []int, counter []int, from int, target int, path []int, ans *[][]int) {

	if target == 0 {
		*ans = append(*ans, append([]int(nil), path...))
		return
	}

	if from == len(uniqNums) {
		return
	}

	if uniqNums[from] <= target {
		//选
		if counter[from] > 0 {
			counter[from]--
			backtracking(uniqNums, counter, from, target - uniqNums[from], append(path, uniqNums[from]), ans)
			counter[from]++
		}
		//不选
		backtracking(uniqNums, counter, from + 1, target, path, ans)
	}
}
// @lc code=end

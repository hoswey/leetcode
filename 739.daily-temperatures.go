/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
import "container/list"

func dailyTemperatures(temperatures []int) []int {

	if len(temperatures) == 0 {
		return nil
	}

	//input: temperatures = [73,74,75,71,69,72,76,73]
	//Output: [1,1,4,2,1,1,0,0]
	S := list.New()
	S.PushFront(0)

	ans := make([]int, len(temperatures))

	for i := 1; i < len(temperatures); i++ {
		temper := temperatures[i]
		for S.Len() != 0 && temperatures[S.Front().Value.(int)] < temper {
			e := S.Front()
			S.Remove(e)
			ans[e.Value.(int)] = i - e.Value.(int)
		}
		S.PushFront(i)
	}
	return ans
}

// @lc code=end


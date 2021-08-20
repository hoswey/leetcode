/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
import "container/list"

func maxSlidingWindow(nums []int, k int) []int {

	l := list.New()

	var ans []int
	for i, num := range nums {

		for l.Len() > 0 && nums[l.Back().Value.(int)] < num {
			l.Remove(l.Back())
		}

		l.PushBack(i)

		for l.Front().Value.(int) <= i - k {
			l.Remove(l.Front())
		}

		if i >= k - 1 {
			ans = append(ans, nums[l.Front().Value.(int)])
		}
	}
	return ans
}
// @lc code=end


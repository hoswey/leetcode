/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
import "container/list"

func findOrder(numCourses int, prerequisites [][]int) []int {

	edges := make([][]int, numCourses)
	//0 未搜索，1 搜索中， 2 已搜索
	visits := make([]int, numCourses)
	S := list.New()
	var invalid bool

	var dfs func(int)
	dfs = func(n int) {

		visits[n] = 1
		for _, k := range edges[n] {

			if invalid {
				return
			}

			if visits[k] == 0 {
				dfs(k)
			} else if visits[k] == 1 {
				invalid = true
				return
			}
		}
		visits[n] = 2
		S.PushFront(n)
	}

	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
	}

	for i, _ := range visits {
		if !invalid && visits[i] == 0 {
			dfs(i)
		}
	}

	if invalid {
		return nil
	}

	var ans []int
	for i := S.Len(); i > 0; i-- {

		e := S.Front()
		S.Remove(e)

		ans = append(ans, e.Value.(int))
	}
	return ans
}

// @lc code=end


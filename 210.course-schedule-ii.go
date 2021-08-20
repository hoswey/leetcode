/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
import "container/list"

func findOrder(numCourses int, prerequisites [][]int) []int {

	edges := make([][]int, numCourses)
	indegs := make([]int, numCourses)

	for _, pre := range prerequisites {

		edges[pre[1]] = append(edges[pre[1]], pre[0])
		indegs[pre[0]]++
	}

	l := list.New()
	for i, indeg := range indegs {
		if indeg == 0 {
			l.PushBack(i)
		}
	}

	var ans []int
	for l.Len() > 0 {

		e := l.Front()
		l.Remove(e)

		course := e.Value.(int)
		ans = append(ans, course)
		
		for _, edge := range edges[course] {
			indegs[edge]--
			if indegs[edge] == 0 {
				l.PushBack(edge)
			}
		}
	}

	if len(ans) != numCourses {
		return nil
	}

	return ans
}

// @lc code=end


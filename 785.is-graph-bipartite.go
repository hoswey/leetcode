/*
 * @lc app=leetcode id=785 lang=golang
 *
 * [785] Is Graph Bipartite?
 */

// @lc code=start
import "container/list"

func isBipartite(graph [][]int) bool {

	if len(graph) == 0 {
		return false
	}

	colors := make([]int, len(graph))
	l := list.New()
	
	for k := 0; k < len(graph); k++ {
		
		l.PushBack(k)
		for l.Len() > 0 {
			for i := l.Len(); i > 0; i-- {
				e := l.Front()
				l.Remove(e)
				node := e.Value.(int)

				if colors[node] != 0 {
					continue
				}

				distColors := make(map[int]bool)
				for _, connect := range graph[node] {
					if colors[connect] == 0 {
						l.PushBack(connect)
					} else {
						distColors[colors[connect]] = true
					}
				}

				if len(distColors) == 2 {
					return false
				}

				if len(distColors) == 0 {
					colors[node] = 1
					continue
				}

				for k, _ := range distColors {
					if k == 2 {
						colors[node] = 1
					} else {
						colors[node] = 2
					}
				}
			}
		}
	}
	return true
}
// @lc code=end
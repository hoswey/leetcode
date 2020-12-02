/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph(node *Node) *Node {

	visited := map[*Node]*Node{}
	var cg func(*Node) *Node
	cg = func(node *Node) *Node {

		if node == nil {
			return nil
		}

		if val, ok := visited[node]; ok {
			return val
		}

		cloneNode := &Node{Val:node.Val, Neighbors: []*Node{}}
		visited[node] = cloneNode

		for _, neighbor := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(neighbor))
		}
		return cloneNode
	}

	return cg(node)
}
// @lc code=end


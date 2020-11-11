/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {

	var levels [][]*Node
	recurse(&levels, 0, root)
	return root
}

func recurse(levels *[][]*Node, l int, node *Node) {

	if node == nil {
		return
	}

	if len(*levels) == l {
		*levels = append(*levels,[]*Node{node})
	} else {
		(*levels)[l][len((*levels)[l])-1].Next = node
		(*levels)[l] = append((*levels)[l], node)
	}

	recurse(levels, l+1, node.Left)
	recurse(levels, l+1, node.Right)
}
// @lc code=end


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

	if root == nil {
		return nil
	}

	var levels [][]*Node
	doConnect(root, &levels, 0)	
	return root
}

func doConnect(root *Node, levels *[][]*Node, level int) {

	if root == nil {
		return
	}

	if len(*levels) == level {
		*levels = append(*levels, []*Node{root})
	} else {
		(*levels)[level][len((*levels)[level]) - 1].Next = root
		(*levels)[level] = append((*levels)[level], root)
	}

	doConnect(root.Left, levels, level + 1)
	doConnect(root.Right, levels, level + 1)	
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    
    return recurseGenerateTrees(1, n)
}

func cloneTree(treeNode *TreeNode) *TreeNode {

	if treeNode == nil {
		return nil
	}

	left := cloneTree(treeNode.Left)
	right := cloneTree(treeNode.Right)

	node := TreeNode{Val: treeNode.Val, Left: left, Right: right}
	return &node
}

func recurseGenerateTrees(start, end int) []*TreeNode {

	if start > end {
		return nil
	}

	if start == end {
		node := TreeNode{Val: start}
		return []*TreeNode{&node}
	}

	var nodes []*TreeNode
	for i := start; i <= end; i++ {

		var curNodes []*TreeNode

		leftSubTrees  := recurseGenerateTrees(start, i - 1)
		rightSubTrees := recurseGenerateTrees(i + 1, end)		
		
		if leftSubTrees == nil {
			root :=  TreeNode{Val: i}
			curNodes = append(curNodes, &root)
		} else {
			for _, l := range leftSubTrees {
				root :=  TreeNode{Val: i}
				root.Left = l
				curNodes = append(curNodes, &root)
			}		
		}	

		if rightSubTrees != nil {
			existNodes := curNodes
			curNodes = nil
			for _, n := range existNodes {
				for _, r := range rightSubTrees {
					root := cloneTree(n)
					root.Right = r
					curNodes = append(curNodes, root)
				}
			}
		}
		nodes = append(nodes, curNodes...)
	}

	return nodes
}
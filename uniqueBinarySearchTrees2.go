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

func recurseGenerateTrees(start, end int) []*TreeNode {

	if start > end {
		return []*TreeNode{nil}
	}

	if start == end {
		node := TreeNode{Val: start}
		return []*TreeNode{&node}
	}

	var nodes []*TreeNode
	for i := start; i <= end; i++ {

		leftSubTrees  := recurseGenerateTrees(start, i - 1)
		rightSubTrees := recurseGenerateTrees(i + 1, end)		
		
		for _, l := range leftSubTrees {
			for _, r := range rightSubTrees {
				root :=  TreeNode{Val: i}
				root.Left = l
				root.Right = r  
                nodes = append(nodes, &root)
			}
		}
	}
	return nodes
}
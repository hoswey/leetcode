/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

	var val []int
	recuse(root, &val)
	return val
}

func recurse(root *TreeNode, vals *[]int) {


	if root == nil {
		return nil
	}

	if root.Left != nil {
		return recurse(root.Left, val)
	}

 	*val = append(*val, root.Val)

 	if root.Right != nil {
		return recurse(root.Right, val)
	}
}
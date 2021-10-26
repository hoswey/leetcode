/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

	return doIsValidBST(root, nil, nil)
}

func doIsValidBST(node *TreeNode, low, upper *int) bool {

	if node == nil {
		return true
	}

	if low != nil && node.Val <= *low || upper != nil && node.Val >= *upper {
		return false
	}

	return doIsValidBST(node.Left, low, &node.Val) && doIsValidBST(node.Right, &node.Val, upper)
}
// @lc code=end


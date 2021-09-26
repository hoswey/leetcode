/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return recurse(root.Left, root.Right)
}

func recurse(n1, n2 *TreeNode) bool {

	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.Val != n2.Val {
		return false
	}

	return recurse(n1.Left, n2.Right) && recurse(n1.Right, n2.Left)
}
// @lc code=end


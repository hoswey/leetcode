/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && abs(getHeight(root.Left) - getHeight(root.Right)) <= 1
}

func getHeight(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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

	return abs(height(root.Left), height(root.Right)) <= 1 && 
		isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {

	if node == nil {
		return 0
	}

	heightOfLeft := height(node.Left)
	heightOfRight := height(node.Right)
	max := heightOfLeft
	if heightOfRight > max {
		max = heightOfRight
	}
	return max + 1
}

func abs(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}
// @lc code=end


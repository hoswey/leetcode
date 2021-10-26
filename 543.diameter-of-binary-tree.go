/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func diameterOfBinaryTree(root *TreeNode) int {

	var ans int
	var getLenOfChild func(*TreeNode) int

	getLenOfChild = func(root *TreeNode) int {

		if root == nil {
			return 0
		}

		left := getLenOfChild(root.Left)
		right := getLenOfChild(root.Right)

		ans = max(ans, left+right+1)
		return max(left, right) + 1
	}
	getLenOfChild(root)

	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


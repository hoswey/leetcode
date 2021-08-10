/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
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
func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}

	var ans int
	dfs(root.Left, true, &ans)
	dfs(root.Right, false, &ans)

	return ans
}

func dfs(node *TreeNode, left bool, ans *int) {

	if node == nil {
		return 
	}

	if node.Left == nil && node.Right == nil {
		if left {
			*ans += node.Val
		}
		return
	}

	dfs(node.Left, true, ans)
	dfs(node.Right, false, ans)
}

// @lc code=end


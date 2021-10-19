/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
func maxPathSum(root *TreeNode) int {
	// 对每个节点进行遍历，结果两种情况
	// 1、路径经过左右子树，那么root包含在其中
	// 2、路径只包含 左/右子树， 那么输出大的值

	ans := -(1 << 31)

	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {

		if root == nil {
			return 0
		}

		leftSum := dfs(root.Left)
		rightSum := dfs(root.Right)

		// fmt.Printf("leftSum: %v, rightSum: %v \n", leftSum, rightSum)

		ans = max(ans, max(leftSum, 0)+max(rightSum, 0)+root.Val)
		return max(max(leftSum, rightSum), 0) + root.Val
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

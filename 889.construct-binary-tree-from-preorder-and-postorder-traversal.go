/*
 * @lc app=leetcode id=889 lang=golang
 *
 * [889] Construct Binary Tree from Preorder and Postorder Traversal
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[0]

	if len(preorder) == 1 {
		return root
	}

	lVal := preorder[1]
	var i int
	for postorder[i] != lVal {
		i++
	}

	l := i + 1
	root.Left = constructFromPrePost(preorder[1: 1 + l], postorder[0: l])
	if len(preorder) > (1 + l) {
		root.Right = constructFromPrePost(preorder[1+l:], postorder[l: len(postorder) - 1])
	}
	return root
}
// @lc code=end


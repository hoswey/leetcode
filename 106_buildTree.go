/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {

	if len(inorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder) - 1]

	var idx int
	for i, v := range inorder {
		if v == rootVal {
			idx = i
			break
		}
	}

	root := new(TreeNode)
	root.Val = rootVal
	root.Left = buildTree(inorder[0:idx], postorder[0:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx: len(postorder) - 1])
	return root
}
// @lc code=end


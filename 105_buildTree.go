/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
 func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	rootValue := preorder[0]

	var idx int
	for i, v := range inorder {
		if v == rootValue {
			idx = i
		}
	}

	root := new(TreeNode)
	root.Val = rootValue
	root.Left = buildTree(preorder[1: idx+1], inorder[0:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return root
}
// @lc code=end


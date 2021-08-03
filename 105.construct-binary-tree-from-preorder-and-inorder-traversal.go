/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

	root := new(TreeNode)
    root.Val = preorder[0]

    var inorderIdx int
    for i, v := range inorder {
        if v == root.Val {
            inorderIdx = i
            break
        }
    }

    root.Left = buildTree(preorder[1:1+inorderIdx], inorder[0: inorderIdx])
    root.Right = buildTree(preorder[1+inorderIdx:], inorder[inorderIdx+1:])

    return root
}
// @lc code=end


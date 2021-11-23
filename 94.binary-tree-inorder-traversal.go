/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
import "container/list"

func inorderTraversal(root *TreeNode) []int {

	l := list.New()
	cur := root

	var ans []int
	for l.Len() > 0 || cur != nil {

		for cur != nil {
			l.PushFront(cur)
			cur = cur.Left
		}

		e := l.Front()
		l.Remove(e)

		node := e.Value.(*TreeNode)
		ans = append(ans, node.Val)

		cur = node.Right
	}
	return ans
}
// @lc code=end


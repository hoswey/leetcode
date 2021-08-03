/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var ans []int
	l := list.New()
	l.PushBack(root)

	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)

		node := e.Value.(*TreeNode)
		ans = append(ans, node.Val)

		if node.Right != nil {
			l.PushFront(node.Right)
		}
		if node.Left != nil {
			l.PushFront(node.Left)
		}
	}
	return ans
}
// @lc code=end
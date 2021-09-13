/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	l := list.New()
	l.PushFront(root)
	var ans [][]int
	for l.Len() > 0 {

		var levels []int
		for i := l.Len(); i > 0; i-- {

			e := l.Front()
			l.Remove(e)

			node := e.Value.(*TreeNode)

			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}

			levels = append(levels, node.Val)
		}
		ans = append(ans, levels)
	}
	return ans
}
// @lc code=end


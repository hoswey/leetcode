/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
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

func findBottomLeftValue(root *TreeNode) int {

	if root == nil {
		return 0
	}

	l := list.New()
	l.PushBack(root)

	var ans int
	for l.Len() > 0 {

		var meetLeft bool
		for i := l.Len(); i > 0; i-- {
			e := l.Front()
			l.Remove(e)
			node := e.Value.(*TreeNode)

			if !meetLeft {
				meetLeft = true
				ans = node.Val
			}

			if node.Left != nil {
				l.PushBack(node.Left)
			}

			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
	}
	return ans
}
// @lc code=end
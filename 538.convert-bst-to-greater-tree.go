/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
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
func convertBST(root *TreeNode) *TreeNode {

	
	l := list.New()
	cur := root

	var nodes []*TreeNode
	for l.Len() > 0 || cur != nil {

		for cur != nil {
			l.PushFront(cur)
			cur = cur.Left
		}

		e := l.Front()
		l.Remove(e)
		cur = e.Value.(*TreeNode)
		nodes = append(nodes, cur)
		cur = cur.Right
	}

	for i := len(nodes) - 2; i >= 0; i-- {
		nodes[i].Val = nodes[i].Val + nodes[i+1].Val
	}    
	return root
}
// @lc code=end


/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
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

func findTarget(root *TreeNode, k int) bool {

	if root == nil {
		return false
	}

	var nodes []*TreeNode

	s := list.New()
	cur := root

	for s.Len() > 0 || cur != nil {

		for cur != nil {
			s.PushFront(cur)
			cur = cur.Left
		}

		e := s.Front()
		s.Remove(e)

		cur = e.Value.(*TreeNode)
		nodes = append(nodes, cur)

		cur = cur.Right
	}

	l := 0
	r := len(nodes) - 1

	for l < r {

		sum := nodes[l].Val + nodes[r].Val
		if sum == k {
			return true
		} else if sum < k {
			l++
		} else {
			r--
		}
	}
	return false
}
// @lc code=end


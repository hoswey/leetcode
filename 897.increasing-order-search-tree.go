/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
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

func increasingBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

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

	ans := new(TreeNode)
	head := ans

	for _, node := range nodes {
		fmt.Printf("%d ", node.Val)
		node.Left = nil
		head.Right = node
		head = node
	}
    return ans.Right
}
// @lc code=end


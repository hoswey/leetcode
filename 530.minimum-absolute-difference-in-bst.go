/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
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

func getMinimumDifference(root *TreeNode) int {

	l := list.New()
	var inorder []*TreeNode
	cur := root

	for l.Len() > 0 || cur  != nil {

		for cur != nil {
			l.PushFront(cur)
			cur = cur.Left
		}

		e := l.Front()
		l.Remove(e)

		cur = e.Value.(*TreeNode)
		inorder = append(inorder, cur)

		cur = cur.Right
	}

	ans := (1 << 31) - 1
	for i := 0; i < len(inorder) - 1; i++ {
		if inorder[i+1].Val - inorder[i].Val < ans {
			ans = inorder[i+1].Val - inorder[i].Val 
		} 
	}
	return ans
}
// @lc code=end


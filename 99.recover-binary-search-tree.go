/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode)  {

	var nums []*TreeNode
	s := list.New()
	cur := root

	for s.Len() > 0 || cur != nil {

		for cur != nil {
			s.PushFront(cur)
			cur = cur.Left
		}

		e := s.Front()
		s.Remove(e)

		node := e.Value.(*TreeNode)
		nums = append(nums, node)
		cur = node.Right
	}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i].Val)
	}

	fst, sec := -1, -1
	for i := 1; i < len(nums) ; i++ {
		if nums[i].Val < nums[i-1].Val {
			sec = i
			if fst == -1 {
				fst = i - 1
			} 
		}
	}
	if sec != 0 {
		nums[fst].Val, nums[sec].Val = nums[sec].Val, nums[fst].Val
	} else {
		nums[fst].Val, nums[fst-1].Val = nums[fst-1].Val, nums[fst].Val
	}
}
// @lc code=end


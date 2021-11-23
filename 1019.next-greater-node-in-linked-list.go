/*
 * @lc app=leetcode id=1019 lang=golang
 *
 * [1019] Next Greater Node In Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/list"

func nextLargerNodes(head *ListNode) []int {

	l := list.New()
	cur := head
	for cur != nil {

		for l.Len() != 0 && l.Back().Value.(*ListNode).Val < cur.Val {

			e := l.Back()
			l.Remove(e)

			e.Value.(*ListNode).Val = cur.Val
		}
		l.PushBack(cur)
		cur = cur.Next
	}

	for l.Front() != nil {

		l.Front().Value.(*ListNode).Val = 0
		l.Remove(l.Front())
	}

	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

// @lc code=end

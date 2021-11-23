/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	var pre *ListNode

	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy

	for i := 0; i < right; i++ {

		if i < left {
			pre = cur
		}

		cur = cur.Next
	}

	l2 := pre.Next
	l2Head := l2

	pre.Next = nil
	l3 := cur.Next
	cur.Next = nil

	pre.Next = reverse(l2)
	l2Head.Next = l3

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {

	var pre *ListNode

	for head != nil {

		temp := head.Next
		head.Next = pre

		pre = head
		head = temp
	}

	return pre
}
// @lc code=end


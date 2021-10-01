/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {

	dummy := new(ListNode)
	dummy.Next = head

	slow, fast := dummy, dummy
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	l1 := head
	l2 := slow.Next
	slow.Next = nil

	l2 = reverse(l2)

	dummy = new(ListNode)
	cur := dummy
	for l1 != nil && l2 != nil {

		cur.Next = l1
		cur = cur.Next
		l1 = l1.Next

		cur.Next = l2
		cur = cur.Next
		l2 = l2.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	head = dummy.Next
}

func reverse(head *ListNode) *ListNode {

	var pre *ListNode

	for head != nil {
		cur := head
		head = head.Next
		cur.Next = pre
		pre = cur
	}

	return pre
}

// @lc code=end


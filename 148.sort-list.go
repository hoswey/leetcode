/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	l, r := partition(head)
	l = mergeSort(l)
	r = mergeSort(r)

	return merge(l, r)
}

func partition(head *ListNode) (*ListNode, *ListNode) {

	dummy := new(ListNode)
	dummy.Next = head

	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	next := slow.Next
	slow.Next = nil
	return dummy.Next, next
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := new(ListNode)
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} 

	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}
// @lc code=end


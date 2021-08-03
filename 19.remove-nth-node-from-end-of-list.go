/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head

	var l int
	p := head

	for p != nil {
		p = p.Next
		l++
	}
	p = dummy
	for i := l - n; i > 0; i-- {
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}
// @lc code=end
/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	oh := head
	eh := head.Next
	cur := eh.Next

	co := oh
	ce := eh

	for cur != nil {

		co.Next = cur
		co = co.Next

		cur = cur.Next

		ce.Next = cur
		ce = ce.Next

		if cur == nil {
			break
		}
		cur = cur.Next
	}

	co.Next = eh
	return oh
}
// @lc code=end


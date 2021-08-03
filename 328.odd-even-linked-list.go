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
	oc := oh

	eh := head.Next
	ec := eh

	cur := eh.Next

	for cur != nil {

		oc.Next = cur
		oc = cur

		cur = cur.Next
		ec.Next = cur
		ec = cur

		if cur == nil {
			break
		}
		cur = cur.Next
	}
	oc.Next = eh
	return oh
}
// @lc code=end


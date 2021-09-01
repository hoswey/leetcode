/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

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


/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	slow := head
	fast := head

	for {
		slow = slow.Next
		fast = fast.Next
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next

		if fast == slow {
			return true
		}
	}
	return true
}

// @lc code=end


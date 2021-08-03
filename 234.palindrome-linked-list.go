/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	dummy := new(ListNode)
	dummy.Next = head

	fast, slow := dummy, dummy

	for fast.Next != nil && slow.Next != nil {
		
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}

		slow = slow.Next
	}

	l2 := slow.Next
	slow.Next = nil
	l2 = reverse(l2)
	l1 := dummy.Next

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {

	var pre *ListNode

	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head 
		head = next
	}
	return pre
}
// @lc code=end


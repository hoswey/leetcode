/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := new(ListNode)
	cur := dummy
	var carry int
	for l1 != nil || l2 != nil {

		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		node := new(ListNode)
		sum := v1 + v2 + carry
		node.Val = sum % 10
		carry = sum / 10

		cur.Next = node
		cur = node
	}

	if carry != 0 {
		node := ListNode{Val: carry}
		cur.Next = &node
	}

	return dummy.Next
}

// @lc code=end


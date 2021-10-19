/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
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

	l1 = reverse(l1)
	l2 = reverse(l2)

	var ans, cur *ListNode
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {

		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		node := new(ListNode)
		node.Val = sum % 10

		if cur == nil {
			cur = node
			ans = cur
		} else {
			cur.Next = node
			cur = cur.Next
		}
	}
	return reverse(ans)
}

func reverse(l *ListNode) *ListNode {
	var pre *ListNode
	for l != nil {
		next := l.Next
		l.Next = pre
		pre = l
		l = next
	}
	return pre
}
// @lc code=end

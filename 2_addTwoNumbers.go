/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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

	d1 := new(ListNode)
	d2 := new(ListNode)

	d1.Next = l1
	d2.Next = l2

	l := new(ListNode)
	d := l
	
	var carry int
	for {

		if d1.Next == nil && d2.Next == nil {
			break
		}

		var v1, v2 int
		if d1.Next == nil {
			d2 = d2.Next
			v2 = d2.Val
		} else if d2.Next == nil {
			d1 = d1.Next
			v1 = d1.Val
		} else {
			d1 = d1.Next
			d2 = d2.Next
			v1 = d1.Val
			v2 = d2.Val
		}

		plus := v1 + v2 + carry
		n := new(ListNode)
		n.Val = plus % 10
		carry = plus / 10

		l.Next = n
		l = l.Next
	}

	if carry > 0 {
		n := ListNode{Val: carry}
		l.Next = &n
	}
	return d.Next
}
// @lc code=end


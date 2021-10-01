/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head

	var pre, fir, sec, next *ListNode

	pre = dummy

	for pre != nil {

		fir = pre.Next
		if fir == nil {
			break
		}

		sec = fir.Next
		if sec == nil {
			break
		}

		next = sec.Next

		pre.Next = sec
		sec.Next = fir
		fir.Next = next

		pre = fir
	}

	return dummy.Next
}

// @lc code=end


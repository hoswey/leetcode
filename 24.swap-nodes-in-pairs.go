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
	pre := dummy
	pre.Next = head

	var fir, sec *ListNode
	for {

		if pre == nil {
			break
		}

		fir = pre.Next
		if fir == nil {
			break
		}

		sec = fir.Next
		if sec == nil {
			break
		}

		temp := sec.Next
		pre.Next = sec
		sec.Next = fir
		fir.Next = temp
		
		pre = fir
	}
	return dummy.Next
}
// @lc code=end


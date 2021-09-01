/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy

	for pre.Next != nil {

		tail := pre.Next
		cur := pre.Next

		var end bool
		for i := 0; i < k-1; i++ {
			if cur.Next == nil {
				end = true
				break
			}
			cur = cur.Next
		}

		if end {
			break
		}

		pre.Next = cur
		remain := cur.Next
		cur.Next = nil
		reverse(tail)
		pre = tail
		tail.Next = remain
	}

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {

	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		if next == nil {
			break
		}
		head = next
	}
	return pre
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	if head == nil {
		return nil
	}

	little := new(ListNode)
	littleCur := little
	large := new(ListNode)
	largeCur := large

	for head != nil {

		temp := head.Next
		if head.Val < x {
			littleCur.Next = head
			littleCur = littleCur.Next
		} else {
			largeCur.Next = head
			largeCur = largeCur.Next
		}
		head.Next = nil
		head = temp
	}
	littleCur.Next = large.Next
	return little.Next
}
// @lc code=end


/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	pre := head
	cur := pre.Next
	for cur != nil {
		if cur.Val != pre.Val {
			pre.Next = cur
			pre = cur
		}
		cur = cur.Next
	}  
	pre.Next = nil
	return head
}
// @lc code=end


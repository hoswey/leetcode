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

	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy

	var next, temp *ListNode
	for {

		temp = pre
		for i := k; i > 0; i-- {
			temp = temp.Next
			if temp == nil {
				goto END
			}
		}

		next = temp.Next
		temp.Next = nil

		list := pre.Next
		pre.Next = nil
		pre.Next = reverse(list)
		list.Next = next

		pre = list
	}

END:
	return dummy.Next

}

func reverse(head *ListNode) *ListNode {

	var pre, cur *ListNode
	for head != nil {

		cur = head.Next
		head.Next = pre
		pre = head

		head = cur
	}
	return pre
}
// @lc code=end
/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {
		return nil
	}

	l, m, r := split(head)

	node := new(TreeNode)
	node.Val = m.Val
	node.Left = sortedListToBST(l)
	node.Right = sortedListToBST(r)

	return node
}

func split(head *ListNode) (*ListNode, *ListNode, *ListNode) {

	if head.Next == nil {
		return nil, head, nil
	}

	pre, slow, fast := head, head, head

	for fast != nil && fast.Next != nil {

		pre = slow

		slow = slow.Next
		fast = fast.Next

		if fast == nil {
			break
		}
		fast = fast.Next
	}

	right := slow.Next
	mid := pre.Next
	mid.Next = nil
	pre.Next = nil

	return head, mid, right
}
// @lc code=end
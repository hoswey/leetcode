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

	left, mid, right := split(head)

	node := new(TreeNode)
	node.Val = mid.Val
	node.Left = sortedListToBST(left)
	node.Right = sortedListToBST(right)

	return node
}

func split(head *ListNode) (*ListNode, *ListNode, *ListNode) {

	if head == nil || head.Next == nil {
		return nil, head, nil
	}

	pre, fast, slow := head, head, head
	for fast != nil && fast.Next != nil {

		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

		pre = slow
		slow = slow.Next
	}

	mid := pre.Next
	pre.Next = nil

	right := mid.Next
	mid.Next = nil

	return head, mid, right
}

// @lc code=end

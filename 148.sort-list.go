/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {

	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {

	if head == nil || head.Next == nil{
		return head
	}

	d := new(ListNode)
	d.Next = head

	fast, slow := d, d
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	l1 := d.Next
	l2 := slow.Next
	slow.Next = nil

	l1 = mergeSort(l1)
	l2 = mergeSort(l2)

	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	
	d := new(ListNode)
	cur := d
	for l1 != nil && l2 != nil {
		
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = cur.Next

			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next

			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}
	
	return d.Next
}
// @lc code=end


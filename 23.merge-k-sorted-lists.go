/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	h := len(lists) / 2

	l1 := mergeKLists(lists[0:h])
	l2 := mergeKLists(lists[h:])

	return doMerge(l1, l2)
}

func doMerge(l1, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

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


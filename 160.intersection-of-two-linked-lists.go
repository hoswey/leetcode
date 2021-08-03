/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	l1 := lenOfList(headA)
	l2 := lenOfList(headB)

	if l1 > l2 {
		for i := 0; i < l1 - l2; i++ {
			headA = headA.Next
		}
	}

	if l2 > l1 {
		for i := 0; i < l2 - l1; i++ {
			headB = headB.Next
		}	
	}

	for headA != nil && headB != nil {

		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func lenOfList(h *ListNode) int {
	var l int
	for h != nil {
		h = h.Next
		l++
	}
	return l
}
// @lc code=end


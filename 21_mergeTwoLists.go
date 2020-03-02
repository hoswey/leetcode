/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	n1 := l1 	
	n2 := l2

	var head, tail *ListNode

	for n1 != nil {

		var cur *ListNode

		if n2 == nil {
			cur = n1
            n1 = n1.Next
		} else {
			if n1.Val < n2.Val {
				cur = n1
				n1 = n1.Next
			} else {
				cur = n2
				n2 = n2.Next
			}
		}

		if head == nil {
			head = cur
		} else {
			tail.Next = cur
		}
		cur.Next = nil
		tail = cur
	}

	if n2 != nil {
		tail.Next = n2
	}

	return head   
}
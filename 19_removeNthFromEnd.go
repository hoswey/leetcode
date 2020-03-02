//给定一个链表: 1->2->3->4->5, 和 n = 2.
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//
// 1->2->3->4->5, 和 n = 1, 输出 1->2->3->4
// 1->2->3->4->5, 和 n = 5, 输出 2->3->4->5
//
// 1			  和 n = 1，输出 nil
// nil            和 0    , 输出 nil
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head

	slow := dummy
	fast := dummy

	for i := 0; i < n + 1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	if slow.Next == head {
		ret := head.Next
		head.Next = nil
		return ret
	}

	last := slow.Next.Next
	slow.Next.Next = nil
	slow.Next = last

	return head
}
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 示例:

// 给定 1->2->3->4, 你应该返回 2->1->4->3.
 
 /**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	var pre, fir, sec, next
	dummy := new(ListNode)
	dummy.Next = head

	pre = dummy
	for {

		if pre == nil {
			break
		}

		fir = pre.Next
		if fir == nil {
			break
		}

		sec = fir.Next
		if sec == nil {
			break
		}

		next = sec.Next

		fir.Next = next
		sec.Next = fir
		pre.Next = sec

		pre = fir
	}
	return dummy.Next
}
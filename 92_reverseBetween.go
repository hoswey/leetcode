// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。

// 示例:

// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL
 
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if head == nil || m < 0 || n < m {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head

	var pre, next *ListNode

	pre = dummy
	for i := 0; i < m - 1; i ++ {

		if pre.Next == nil {
			return head
		}
		pre = pre.Next
	}

	if pre.Next == nil {
		return head
	}

	start := pre.Next
	end := start
	for i := 0; i < (n-m); i++ {
		if end.Next == nil {
			break
		}
		end = end.Next
	}

	next = end.Next
	end.Next = nil

	end = start
	start = reverse(start)

	pre.Next = start
	end.Next = next

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var pre *ListNode
	cur := head 

	for {

		temp := cur.Next
		cur.Next = pre
		pre = cur

		cur = temp 

		if cur == nil {
			return pre
		}
	}
}
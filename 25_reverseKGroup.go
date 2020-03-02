/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {

    if head == nil || head.Next == nil {
        return head
    }

	ans := new(ListNode)
	ans.Next = head

	var pre, next, start, end *ListNode

	pre = ans
	start = head
	end = head

	for end != nil {

		for i := 0; i < k - 1; i++ {
			end = end.Next 
			if end == nil {
				break
			}
		}

		if end == nil {
			break
		}

		next = end.Next
		end.Next = nil

		pre.Next = reverse(start)
		start.Next = next

		pre = start

		start = pre.Next
		end = pre.Next
        
        printList("end", ans.Next)
	}

	return ans.Next
}

func reverse(head *ListNode) *ListNode {

	var pre *ListNode 
	cur := head

    
    printList("reverse 1", head)

	for cur != nil {

		temp := cur.Next
		cur.Next = pre

		pre = cur
		cur = temp
	}

    printList("reverse 2", pre)

	return pre
}

func printList(msg string, h *ListNode) {
    
    fmt.Printf("%s:", msg)
    for h != nil {
        fmt.Printf("%d\t",h.Val)
        h = h.Next
    }
    fmt.Println()
}
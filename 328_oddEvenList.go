/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	odd := new(ListNode)
	even := new(ListNode)

    oddStart := odd
    evenStart := even

	for head != nil {

		odd.Next = head
		if head.Next == nil {
		    odd = odd.Next            
			break
		}
		even.Next = head.Next
		head = head.Next.Next

		odd = odd.Next
		odd.Next = nil
		even = even.Next
		even.Next = nil
	}
	odd.Next = evenStart.Next

	return oddStart.Next
}
func sortList(head *ListNode) *ListNode {

	length := lenOfList(head)
	if length <= 1 {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head

	for seq := 1; seq < length; seq += seq {

		pre := dummy

		for {

			var h1, e1, h2, e2, next *ListNode

			h1 = pre.Next
			e1 = h1

			var eof bool
			for i := 0; i < seq - 1; i++ {
				if e1 == nil {
					break
				}
				e1 = e1.Next
			}


			if e1 == nil || e1.Next == nil{
				break
			}

			h2 = e1.Next
			e1.Next = nil

			e2 = h2
			for i := 0; i < seq - 1; i++ {
				if e2 == nil {
					break
				}
				e2 = e2.Next
			}

			if e2 == nil || e2.Next == nil {
				eof = true
			} else {
				next = e2.Next
				e2.Next = nil
			}

			h, e := merge(h1, h2)
			pre.Next = h
			e.Next = next

			pre = e
			if eof {
				break
			}
		}
	}

    return dummy.Next
}

func merge(h1, h2 *ListNode) (*ListNode, *ListNode) {

	dummy := new(ListNode)

	cur := dummy

	for h1 != nil && h2 != nil {
		var temp *ListNode
		if h1.Val < h2.Val {
			temp = h1
			h1 = h1.Next

		} else {
			temp = h2
			h2 = h2.Next
		}

		cur.Next = temp
		cur = temp
		temp.Next = nil
	}

	if h1 != nil {
		cur.Next = h1
	}

	if h2 != nil {
		cur.Next = h2
	}

	end := cur
	for end.Next != nil {
		end = end.Next
	}

	return dummy.Next, end
}

func lenOfList(head *ListNode) int {

	i := 0
	for head != nil {
		i++
		head = head.Next
	}

	return i
}
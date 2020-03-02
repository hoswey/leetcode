/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 给定这个链表：1->2->3->4->5

// 当 k = 2 时，应当返回: 2->1->4->3->5

// 当 k = 3 时，应当返回: 3->2->1->4->5

func reverseKGroup(head *ListNode, k int) *ListNode {
    
    if head == nil {
        return head
    }

	var ans, tail, left, right *ListNode
	first := true
	
	for {
		// 1. left,right指向tail.Next, 假如tail为nil,则指向第一个指针
		if tail == nil {
			left = head
			right = head
		} else {
            
			left = tail.Next
			right = tail.Next
		}


        fmt.Printf("1 left=%d\n", left.Val)

		// 2. right指针Next k - 1 到达反转的元素
		for i := 0; i < k - 1; i++ {

			if right == nil {
				break
			}
			right = right.Next
		}

		// 	假如是第一次转
		// 		假如right指针不为nil, 那么表示当前元素数量满足k
		// 			那么ans指针指向right指针
		// 		否则
		// 			直接返回head
		// 	否则
		// 		假如right指针为nil, 表示元素不足
		// 			返回ans
		if first {
			if right != nil {
				ans = right
			} else {
				return head
			}
			first = false
		} else {
			if right == nil {
				return ans
			}
		}

		// 	开始反转
		pre := right.Next
		cur := left
        fmt.Printf("2.1, right.value=%d\n", right.Val)

		for {

			temp := cur
			cur = cur.Next

			temp.Next = pre
			if temp == right {
				break
			}
			pre = temp
            fmt.Printf("2.2, cur.value=%d\n", cur.Val)        
			
		}

		// 指向cur的最后一个，移动tail到下一个，cur指向tail
		if tail != nil {
			tail.Next = right
		}

		tail = left
		//假如right的Next为nil,break
		if tail.Next == nil {
			break
		}
	}
	return ans
}
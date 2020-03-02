import "fmt"
// 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

// 示例:

// 输入:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 输出: 1->1->2->3->4->4->5->6

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
    
    if len(lists) < 2 {
        return lists[0]
    }

    var ret *ListNode
	for i := 0; i < len(lists); i++ {
        
        if lists[i] == nil {
            continue
        }
		ret = doMergeTwoLists(ret, lists[i])
		printLists(ret)
	}

	return ret
}

func doMergeTwoLists(first, second *ListNode) *ListNode {

	if first == nil && second != nil {
		return second
	} else if second == nil && first != nil {
		return first
	} else if first == nil && second == nil {
		panic("bug")
	}

	dummy := new(ListNode)
    cur := dummy

	for first != nil && second != nil {
        
		var node *ListNode
		if first.Val <= second.Val {
			node = first
		    first = first.Next            
		} else {
			node = second
	    	second = second.Next            
		}

		cur.Next = node
		cur = cur.Next
		cur.Next = nil
	}

    printLists(first)
	if first != nil {
		cur.Next = first
	}

	if second != nil {
		cur.Next = second
	}

	ans := dummy.Next
	dummy.Next = nil //free the memory

	return ans
}

func printLists(list *ListNode)  {

	// for list != nil {
	// 	list = list.Next
	// }
	// fmt.Println()
}
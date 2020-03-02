/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    
    var l []int
    recurseAppend(head, &l)
    return l
}

func recurseAppend(head *ListNode, l *[]int) {
    
    if head == nil {
        return
    }
    
    if head.Next != nil {
        recurseAppend(head.Next, l)
    }
    
    *l = append(*l, head.Val)
}
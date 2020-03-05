/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
    
    dummy := new(ListNode)
    dummy.Next = head
    
    pre := dummy
    cur := head
    
    for cur != nil && cur.Val != val {        
        cur = cur.Next
        pre = pre.Next        
    }
    
    pre.Next = cur.Next
    return dummy.Next
}
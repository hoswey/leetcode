/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    
    if head == nil {
        return nil
    }
    
    cur := head
    var pre *ListNode
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
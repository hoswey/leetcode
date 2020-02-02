/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    head := new(ListNode)
    headSet := false
    cur := head
    
    var carry int
    for ; l1 != nil; l1 = l1.Next {        
        
        sum := l1.Val + carry
        carry = 0
        
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        
        if sum >= 10 {
            carry = sum / 10            
            sum = sum % 10
        }
        
        if !headSet {
            cur.Val = sum
            headSet = true
        } else {
            newNode := ListNode{Val:sum}
            cur.Next = &newNode
            cur = &newNode
        }
    }
    
    for ; l2 != nil; l2 = l2.Next {
        
        sum := l2.Val + carry
        carry = 0            
        
        if sum >= 10 {
            carry = sum / 10            
            sum = sum % 10
        }

        if !headSet {
            cur.Val = sum
            headSet = true
        } else {
            newNode := ListNode{Val:sum}
            cur.Next = &newNode
            cur = &newNode
        }
    }
    
    if carry != 0 {
        newNode := ListNode{Val:carry}
        cur.Next = &newNode
        cur = &newNode        
    }
    
    return head
}
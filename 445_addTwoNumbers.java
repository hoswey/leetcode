/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        
        Stack<Integer> s1 = new Stack<Integer>();
        Stack<Integer> s2 = new Stack<Integer>();
        
        putToStack(s1, l1);
        putToStack(s2, l2);
        
        ListNode dummy = new ListNode(0);
        
        int carry = 0;
        while (!s1.empty() && !s2.empty()) {
            
            int sum = s1.pop() + s2.pop() + carry;
            carry = sum / 10;                        
            insertNode(dummy,sum%10);
        }
        
        Stack<Integer> remain = null;
        if (!s1.empty()) {
            remain = s1;
        } else if (!s2.empty()) {
            remain = s2;
        }
        
        while (remain != null && !remain.empty()) {
            
            int sum = remain.pop() + carry;
            carry = sum / 10;                        
            insertNode(dummy,sum%10);
        } 
        
        if (carry != 0) {
            insertNode(dummy, carry);
        }
             
        return dummy.next;
    }
        
    public void insertNode(ListNode dummy,int val) {
           
        ListNode node = new ListNode(val);            
        node.next = dummy.next;
        dummy.next = node;                
    }
    
    public void putToStack(Stack<Integer> s, ListNode l) {
        
        while (l != null) {            
            s.push(l.val);
            l = l.next;
        }
    }
}
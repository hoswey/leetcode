/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {

    public ListNode detectCycle(ListNode head) {
        
        if (head == null) {
            return null;
        }
        
        //1. 判断是否有环
        ListNode slow = head;
        ListNode fast = head;
        
        int lenOfCycle = 0;
        while(true) {
    
            if (fast.next == null || fast.next.next == null) {
                return null;
            }
            
            slow = slow.next;
            fast = fast.next.next;
            
            if (slow == fast) {
                //meet         
                //2. 计算环的长度                
                ListNode temp = slow;
                do {                
                    slow = slow.next;
                    lenOfCycle++;                    
                } while(temp != slow); 
                break;
            }
        }
        
        //3. 快慢指针找环
        slow = head;
        fast = head;        
        for (int i = 0; i< lenOfCycle; i++) {
            fast = fast.next;
        }        
        
        while(slow != fast) {
            slow = slow.next;
            fast = fast.next;
        }
    
        return slow;
    }
}
/
*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

  if head == nil {
    return head
  }

  dummy := new(ListNode)
  dummy.Next = head
  slow := dummy
  fast := dummy

  for slow.Next != nil && fast.Next != nil && fast.Next.Next !=nil {

    slow = slow.Next
    fast = fast.Next.Next

    if slow == fast {

      fast = dummy
      for fast != slow {
        fast = fast.Next
        slow = slow.Next
      }
      return slow
    }
  }
  
  return nil
}
// @lc code=end


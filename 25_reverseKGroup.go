/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import (
  "fmt"
)

func reverseKGroup(head *ListNode, k int) *ListNode {

  if head == nil || k <= 1 {
    return head
  }

  dummy := new(ListNode)
  dummy.Next = head

  pre := dummy

  var start, end, next *ListNode
  for pre.Next != nil {

    start = pre.Next
    end = pre.Next
    for i:= 0; i < k - 1; i++ {

      if end.Next == nil {
        goto END
      }
      end = end.Next
    }

    next = end.Next
    end.Next = nil

    pre.Next = reverse(start)
    start.Next = next

    pre = start
  }

  END:
  return dummy.Next
}

func reverse(head *ListNode) *ListNode {

  if head == nil {
    return nil
  }

  var pre *ListNode
  cur := head

  for {

    temp := cur.Next
    cur.Next = pre
    pre = cur

    if temp == nil {
      return pre
    }
    cur = temp
  }
}

func printList(head *ListNode) []int {

  var ret []int
  for head != nil {
    ret = append(ret, head.Val)
    head = head.Next
  }

  return ret
}

type ListNode struct {
  Val int
  Next *ListNode
}

func createList(arr []int) *ListNode {

  dummy := new(ListNode)
  h := dummy
  for i := range arr {
    n := new(ListNode)
    n.Val = arr[i]

    h.Next = n
    h = h.Next
  }
  return dummy.Next
}

func main() {

  var input []int
  var k int

  input = []int{1,2,3,4,5}
  k = 2

  fmt.Printf("input:%v, output:%v, expected:%v\n", input, printList(reverseKGroup(createList(input), k)), []int{2,1,4,3,5})

  k = 3
  fmt.Printf("input:%v, output:%v, expected:%v\n", input, printList(reverseKGroup(createList(input), k)), []int{3,2,1,4,5})
}


package main

import "fmt"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 示例 1:
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.

// 示例 2:
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

func reorderList(head *ListNode)  {

	//拆分链表， p1为源链表的前半部分， p2为后半部分的逆转链表
	var l int
	temp := head
	for temp != nil {
		temp = temp.Next
		l++
	}

	if l <= 2 {
		return
	}

	p1 := head
	lastNode := head

	l1 := (l+1)/2
	for i := 0; i < (l1 - 1); i++ {
		lastNode = lastNode.Next
	}

	p2 := lastNode.Next
	lastNode.Next = nil

	p2 = reverse(p2)
	
	//合并链表
	for p1 != nil && p2 != nil {
        temp := p2 
        p2 = p2.Next

        temp.Next = p1.Next
        p1.Next = temp 
        p1 = temp.Next     
	}
}

func reverse(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head 

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



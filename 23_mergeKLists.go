package main

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

    lists = filterNilList(lists)
    return merge(lists)

}

func merge(lists []*ListNode) *ListNode {

    if len(lists) == 0 {
        return nil
    }

    if len(lists) == 1 {
        return lists[0]
    }

    mid := len(lists)/2

    a := merge(lists[0:mid])
    b := merge(lists[mid: len(lists)])

	return doMergeTwoLists(a, b)
}

func doMergeTwoLists(a, b *ListNode) *ListNode{

	dummy := new(ListNode)
	cur := dummy
	for a != nil && b != nil {

		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
			cur = cur.Next
		} else {
			cur.Next = b
			b = b.Next
			cur = cur.Next
		}
	}

	if a != nil {
		cur.Next = a
	}

	if b != nil {
		cur.Next = b
	}
	return dummy.Next
}

func filterNilList(lists []*ListNode) []*ListNode {

	var temp  []*ListNode
	for i, _ := range lists {
		if lists[i] != nil {
			temp = append(temp, lists[i])
		}
	}
	return temp
}

func main() {

	a := buildList([]int{1,4,5})
	b := buildList([]int{1,3,4})
	c := buildList([]int{2,6})

	printListNode(mergeKLists([]*ListNode{a,b,c}))
}

func printListNode(list *ListNode) {

	for list != nil {
		fmt.Printf("%d\t",list.Val)
		list = list.Next
	}
}

func buildList(arr []int) *ListNode {

	dummy := new(ListNode)
	cur := dummy
	for i, _ := range arr {
		node := ListNode{Val: arr[i]}
		cur.Next = &node
		cur = cur.Next
	}

	return dummy.Next
}



type ListNode struct {
  Val int
  Next *ListNode
 }

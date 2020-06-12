
// 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

// 示例:

// 给定的有序链表： [-10, -3, 0, 5, 9],

// 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

//       0
//      / \
//    -3   9
//    /   /
//  -10  5


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	fast := head
	slow := head
	pre := slow

	for fast !=nil && fast.Next != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		pre = slow
		slow = slow.Next
	}

	root := TreeNode{Val: slow.Val}
	rightList := slow.Next
	slow.Next = nil

	pre.Next = nil

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(rightList)

	return &root
}
/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"

type Stack struct {
	l *list.List
}

func NewStack() *Stack {

	s := new(Stack)
	s.l = list.New()

	return s
}

func (this *Stack) Pop() *TreeNode {

	if this.l.Len() == 0 {
		return nil
	}

	e := this.l.Front();
	this.l.Remove(e)
	return e.Value.(*TreeNode)
}

func (this *Stack) IsEmpty() bool {
	return this.l.Len() == 0
}

func (this *Stack) Push(node *TreeNode) {
	this.l.PushFront(node)
}

func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	s1 := NewStack()
	s2 := NewStack()

	s1.Push(root)

	for !s1.IsEmpty() {

		node := s1.Pop()
		s2.Push(node)
		if node.Left != nil {
			s1.Push(node.Left)
		}
		if node.Right != nil {
			s1.Push(node.Right)
		}
	}

	var ans []int
	for !s2.IsEmpty() {
		ans = append(ans, s2.Pop().Val)
	}
	return ans
}
// @lc code=end


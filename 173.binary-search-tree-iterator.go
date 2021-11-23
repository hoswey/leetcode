/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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

type BSTIterator struct {
	l   *list.List
	cur *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {

	return BSTIterator{cur: root, l: list.New()}
}

func (this *BSTIterator) Next() int {

	for this.cur != nil {
		this.l.PushBack(this.cur)
		this.cur = this.cur.Left
	}

	e := this.l.Back()
	this.l.Remove(e)

	val := e.Value.(*TreeNode).Val
	this.cur = e.Value.(*TreeNode).Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || this.l.Len() != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end


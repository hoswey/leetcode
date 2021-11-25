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
	return BSTIterator{l: list.New(), cur: root}
}

func (this *BSTIterator) Next() int {

	for this.cur != nil {
		this.l.PushFront(this.cur)
		this.cur = this.cur.Left
	}

	node := this.l.Front()
	this.l.Remove(node)

	this.cur = node.Value.(*TreeNode).Right

	return node.Value.(*TreeNode).Val
}

func (this *BSTIterator) HasNext() bool {
	return this.l.Len() > 0 || this.cur != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end


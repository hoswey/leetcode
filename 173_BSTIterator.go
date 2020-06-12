/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	idx int
	vals []int
}


func Constructor(root *TreeNode) BSTIterator {

	var vals []int
    inOrder(root, &vals)    
	bst := BSTIterator{vals: vals , idx: -1}
	return bst
}

func inOrder(node *TreeNode, vals *[]int) {

	if node == nil {
		return
	}

	inOrder(node.Left, vals)
	*vals = append(*vals, node.Val)
	inOrder(node.Right, vals)
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {

	this.idx++
	return this.vals[this.idx]
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {

	return len(this.vals) > (this.idx + 1)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
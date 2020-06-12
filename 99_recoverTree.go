/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {

	if root == nil {
		return
	}

	var arr []*TreeNode
	inOrder(root, &arr)

	var reverse []*TreeNode

	for i := 0; i < len(arr) - 1; i++ {
		if arr[i].Val > arr[i+1].Val{
			reverse = append(reverse, arr[i])
			reverse = append(reverse, arr[i+1])
		}
	}

	reverse[0].Val, reverse[len(reverse) - 1].Val = reverse[len(reverse) - 1].Val, reverse[0].Val
}

func inOrder(node *TreeNode, arr *[]*TreeNode) {

	if node == nil {
		return
	}

	inOrder(node.Left, arr)
	*arr = append(*arr, node)
	inOrder(node.Right, arr)
}
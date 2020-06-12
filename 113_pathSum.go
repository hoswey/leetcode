 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {

	if root == nil {
		return nil
	}

	var ans [][]int
	recurse(nil, root, sum, &ans)
	return ans
}

func recurse(path []*TreeNode, node *TreeNode, target int, ans *[][]int) {

	if node.Left == nil && node.Right == nil {

		if target == node.Val {
			var vals []int
			for i := range path {
				vals = append(vals, path[i].Val)
			}
			vals = append(vals, node.Val)
			*ans = append(*ans, vals)
		}
		return
	}

	path = append(path, node)

	if node.Left != nil {
		recurse(path, node.Left, target - node.Val, ans)
	}

	if node.Right != nil {
		recurse(path, node.Right, target - node.Val, ans)
	}
}
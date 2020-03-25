/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    
    if root == nil {
        return nil
    }

    var ans [][]int
    travel(&ans, root, 0)

    return ans
}

func travel(ans *[][]int, root *TreeNode, level int) {

	if root == nil {
		return
	}

	if len(*ans) == level {
		var levelArr []int
		*ans = append(*ans, levelArr)
	}


	if level % 2 == 0 {
		(*ans)[level] = append((*ans)[level], root.Val)
	} else {
        (*ans)[level] = append([]int{root.Val}, (*ans)[level]...)
	}

	travel(ans, root.Left, level + 1)
	travel(ans, root.Right, level + 1)
}
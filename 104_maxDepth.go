/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

    if root == nil {
        return 0
    }

    depth := max(maxDepth(root.Left), maxDepth(root.Right)) + 1
    return depth
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

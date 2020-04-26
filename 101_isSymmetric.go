/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

    if root == nil {
        return true
    }

    return isSymmetric1(root.Left, root.Right)
}

func isSymmetric1(a, b *TreeNode) bool {

    if a == nil && b == nil {
        return true
    }

    if a == nil || b == nil {
        return false
    }

    if a.Val != b.Val {
        return false
    }

    return isSymmetric1(a.Right, b.Left) && isSymmetric1(a.Left, b.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

    if root == nil {
        return true 
    }

    return recurse(root, nil, nil)
}

func recurse(root *TreeNode, lower, upper *int) bool {

    if lower != nil && root.Val <= *lower {
        return false
    } 

    if upper != nil && root.Val >= *upper {
        return false
    }

    if root.Left != nil && !recurse(root.Left, lower, &root.Val) {
        return false
    }

    if root.Right != nil && !recurse(root.Right, &root.Val, upper) {
        return false
    }

    return true
}
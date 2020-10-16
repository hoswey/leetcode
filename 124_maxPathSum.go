/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxPathSum(root *TreeNode) int {

    ans := -1<<31
    recurse(root, &ans)
    return ans
}

func recurse(root *TreeNode, ans *int) int{

    if root == nil {
        return 0
    }

    left := max(recurse(root.Left, ans), 0)
    right := max(recurse(root.Right, ans), 0)

    *ans = max(*ans, left + right + root.Val)

    return root.Val + max(left, right)

}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}


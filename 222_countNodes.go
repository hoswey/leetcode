/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    var sizes []int
    doCountNodes(root, &sizes, 0)
    
    var ans int
    for _, v := range sizes {
        ans += v
    }
    return ans
}

func doCountNodes(node *TreeNode, sizes *[]int, level int) {

    if node == nil {
        return
    }
    if len(*sizes) == level{
        *sizes = append(*sizes, 0)
    }

    (*sizes)[level]++
    doCountNodes(node.Left,  sizes, level+1)
    doCountNodes(node.Right, sizes, level+1)
}
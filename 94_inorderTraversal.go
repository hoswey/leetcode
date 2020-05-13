/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

  var ans []int
  travel(&ans, root)
  return ans
}

func travel(ans *[]int, node *TreeNode) {

  if node == nil {
    return
  }

  *ans = append(*ans, node.Val)
  travel(ans, node.Left)
  travel(ans, node.Right)
}

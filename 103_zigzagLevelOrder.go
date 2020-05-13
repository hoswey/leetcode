/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {

  var ans [][]int
  travel(&ans, root, 0)
  return ans
}

func travel(ans *[][]int, node *TreeNode, level int) {

  if node == nil {
    return
  }

  if len(*ans) == level {
    var level []int
    *ans = append(*ans, level)
  }

  if level % 2 == 0 {
    (*ans)[level] = append((*ans)[level], node.Val)
  } else {
    (*ans)[level] = append([]int{node.Val}, (*ans)[level]...)
  }

  travel(ans, node.Left, level+1)
  travel(ans, node.Right, level+1)
}

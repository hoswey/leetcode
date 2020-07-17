// 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

//  

// 示例 :
// 给定二叉树

//           1
//          / \
//         2   3
//        / \
//       4   5
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func diameterOfBinaryTree(root *TreeNode) int {

  if root == nil {
    return 0
  }

  var ans int
  dfs(root, &ans)
  return ans - 1
}

func dfs(root *TreeNode, ans *int) int {

  if root == nil {
    return 0
  }

  left = dfs(root.Left)
  right = dfs(root.Right)

  *ans = max(*ans, left + right + 1)

  return max(left, right)
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}


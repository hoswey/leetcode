/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	dfs("", root, &ans)
	return ans
}

func dfs(path string, node *TreeNode, ans *[]string) {

	if node == nil {
		return
	}

	if path != "" {
		path += "->"
	} 

	path += strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		*ans = append(*ans, path)
		return
	}

	dfs(path, node.Left, ans)
	dfs(path, node.Right, ans)
}
// @lc code=end


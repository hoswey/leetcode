/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
import "strings"
import "strconv"

func binaryTreePaths(root *TreeNode) []string {


	var allPaths [][]int

	var dfs func(*TreeNode, []int)
	dfs = func(node *TreeNode, path []int) {

		if node == nil {
			return
		}

		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil {
			allPaths = append(allPaths, append([]int(nil), path...))
			return
		}

		if node.Left != nil {
			dfs(node.Left, path)
		} 
		if node.Right != nil {
			dfs(node.Right, path)
		}
	}

	dfs(root, nil)

	var ans []string
	for i := 0; i < len(allPaths); i++ {

		var pathStr []string
		for _, v := range allPaths[i] {
			pathStr = append(pathStr, strconv.Itoa(v))
		}

		ans = append(ans, strings.Join(pathStr, "->"))
	}
	return ans
}
// @lc code=end


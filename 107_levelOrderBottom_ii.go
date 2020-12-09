/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {

	var ans [][]int
	recurse(root, 0, &ans)

	i := 0
	j := len(ans) - 1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}

func recurse(node *TreeNode, level int, ans *[][]int) {

	if node == nil {
		return
	}

	if len(*ans) == level {
		*ans = append(*ans, []int{})
	} 

	(*ans)[level] = append((*ans)[level], node.Val)

	recurse(node.Left, level+1, ans)
	recurse(node.Right, level+1, ans)
} 

// @lc code=end


/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
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
import "container/list"
func averageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return nil
	}

	var levels [][]int
	l := list.New()
	l.PushBack(root)


	for l.Len() > 0 {

		var nums []int

		for i := l.Len(); i > 0; i-- {
			e := l.Front()
			l.Remove(e)

			node := e.Value.(*TreeNode)
			nums = append(nums, node.Val)

			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		levels = append(levels, nums)
	}

	var ans []float64
	for _, level := range levels {

		var sum float64
		for _, l := range level {
			sum += float64(l)
		}
		ans = append(ans, sum/float64(len(level)))
	}
	return ans
}
// @lc code=end


// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// 示例:

// 输入: [1,2,3,null,5,null,4]
// 输出: [1, 3, 4]
// 解释:

//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---


// 链接：https://leetcode-cn.com/problems/binary-tree-right-side-view

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var arr [][]int
	travel(&arr, root, 0)

	var ans []int
	for _, levelArr := range arr {
		ans = append(ans, levelArr[len(levelArr) - 1])
	}

	return ans
}


func travel(arr *[][]int, node *TreeNode, level int) {

	if node == nil {
		return
	}

	if len(*arr) == level {
		var levelArr []int
		*arr = append(*arr, levelArr)
	}

	(*arr)[level] = append((*arr)[level], node.Val)

	travel(arr, node.Left, level + 1)
	travel(arr, node.Right, level + 1)
}
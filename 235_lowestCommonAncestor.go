/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p == q {
		return p
	}

	if p == nil || q == nil {
		return nil
	}

	parents := make(map[*TreeNode]*TreeNode)
	buildParents(root, parents)
	
	fmt.Printf("root parent:%v\n", parents[root])

	pHeight := getHeight(p, parents)
	qHeight := getHeight(q, parents)

	pp := p
	qp := q

	if pHeight > qHeight {
		diff := pHeight - qHeight
		for diff > 0 {
			pp = parents[pp]
			diff--
		}
	}

	if qHeight > pHeight {
		diff := qHeight - pHeight
		for diff > 0 {
			qp = parents[qp]
			diff--
		}
	}

	for pp != qp {
		pp = parents[pp]
		qp = parents[qp]
	}

	return pp
}

func getHeight(node *TreeNode, parents map[*TreeNode]*TreeNode) int {

	var height int
	parent, ok := parents[node]
	for ok {
		height++
		node = parent
		parent, ok = parents[node]
	}
	return height
}

func buildParents(root *TreeNode, parents map[*TreeNode]*TreeNode) {

	if root == nil {
		return
	}

	if root.Left != nil {
		parents[root.Left] = root 
		buildParents(root.Left, parents)
	}

	if root.Right != nil {
		parents[root.Right] = root
		buildParents(root.Right, parents)
	}
}
// @lc code=end


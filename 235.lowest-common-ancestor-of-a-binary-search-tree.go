/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
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
import "container/list"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	
	l := list.New()
	l.PushBack(root)

	for l.Len() > 0 {

		for i := l.Len(); i > 0; i-- {
			
			e := l.Front()
			l.Remove(e)

			node := e.Value.(*TreeNode)

			if hasAncestor(node, p, q) && !hasAncestor(node.Left, p, q) && !hasAncestor(node.Right, p, q) {
				return node
			}

			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
	}
	return nil
}

func hasAncestor(root, p, q *TreeNode) bool {
	return hasOneAncesstor(root, p) && hasOneAncesstor(root, q)
}

func hasOneAncesstor(root, p *TreeNode) bool {

	if root == nil {
		return false
	}

	if root == p {
		return true
	}

	return hasOneAncesstor(root.Left, p) || hasOneAncesstor(root.Right, p)
}
// @lc code=end


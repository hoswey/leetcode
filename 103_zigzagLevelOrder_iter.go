import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {

    if root == nil {
        return nil
    }

    var ans[][]int
    level := list.New()
    level.PushBack(root)

    i := 1
    for level.Len() != 0 {

        nextLevel := list.New()
        var vals []int
        for level.Len() != 0 {

            ele := level.Front()
            level.Remove(ele)

            node := ele.Value.(*TreeNode)
            vals = append(vals, node.Val)

            if i % 2 == 0 {
                if node.Right != nil {
                    nextLevel.PushFront(node.Right)
                }
                if node.Left != nil {
                    nextLevel.PushFront(node.Left)
                }
            } else {
                if node.Left != nil {
                    nextLevel.PushFront(node.Left)
                }
                if node.Right != nil {
                    nextLevel.PushFront(node.Right)
                }
            }
        }

        level = nextLevel
        ans = append(ans, vals)
        i++
    }
    return ans
}

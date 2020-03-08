/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {

    if root == nil {
        return nil
    }
    
    var path []int
    var ans [][]int
    
    pathSumRecurse(root, sum, &path, &ans)    
    return ans
}

func pathSumRecurse(node *TreeNode, trg int, path *[]int, ans *[][]int) {
    
    if node.Left == nil && node.Right == nil{        
        if node.Val == trg {                        
            temp := make([]int, len(*path))
            copy(temp, *path)
            temp = append(temp, node.Val)            
            *ans = append(*ans, temp)
        }        
        return
    }
    
    if node.Left != nil {        
        *path = append(*path, node.Val)
        pathSumRecurse(node.Left, trg - node.Val, path, ans)
        *path = (*path)[0:len(*path)-1]
    }

    if node.Right != nil {        
        *path = append(*path, node.Val)
        pathSumRecurse(node.Right, trg - node.Val, path, ans)
        *path = (*path)[0:len(*path)-1]
    }
}
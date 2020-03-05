/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    
    if A == nil || B == nil {
        return false
    }
    
    if isSubStructureRecurse(A, B) {
        return true
    } else {        
        if A.Left != nil && isSubStructure(A.Left, B) {
            return true
        }
        
        if A.Right != nil && isSubStructure(A.Right, B) {
            return true
        }
    }
    
    return false
}

func isSubStructureRecurse(A *TreeNode, B *TreeNode) bool {
    
    if A == nil {
        return false
    }
    
    if A.Val == B.Val {                
        return (B.Left == nil || isSubStructureRecurse(A.Left, B.Left)) && 
            (B.Right == nil || isSubStructureRecurse(A.Right, B.Right))
    } 
    
    return false
}


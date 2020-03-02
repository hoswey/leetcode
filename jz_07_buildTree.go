/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

    if len(preorder) == 0 {
        return nil
    }
    
    root := new(TreeNode)
    root.Val = preorder[0]
    
    //从中序遍历中找出左子树和右子树的范围    
    i := 0
    for inorder[i] != root.Val{
        i++
    }
    
    lenOfLeft := i 
    lenOfRight := len(inorder) - i - 1 
    
    if lenOfLeft != 0 {
        root.Left = buildTree(preorder[1: 1 + lenOfLeft], inorder[0:lenOfLeft])
    }
    
    if lenOfRight != 0 {
        root.Right = buildTree(preorder[1 + lenOfLeft:], inorder[lenOfLeft+1:]) 
    }   
    return root
}
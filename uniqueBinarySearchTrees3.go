//     ,[1,null,3]], [1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return nil
	}
    
 	dp := make([][]*TreeNode, n + 1)
 	node := TreeNode{Val: 1} 

 	dp[1] = []*TreeNode{&node}

	for i := 2; i <= n; i++ {

        var nodes []*TreeNode
		
		for _, pre := range dp[i-1] {

			cloneNode, _ := cloneTreeNode(pre, nil)
			nNode := TreeNode{Val: i, Left: cloneNode}
			nodes = append(nodes, &nNode)

            trg := pre
			for {
                
                if trg == nil {
					break
				}

                cloneNode, trgP := cloneTreeNode(pre, trg)

				nNode := TreeNode{Val:i}
                nNode.Left = trgP.Right
				
				trgP.Right = &nNode

				nodes = append(nodes, cloneNode)
				trg = trg.Right
			}
		}

		dp[i] = nodes
	}
	return dp[n]
}

func cloneTreeNode(node, p *TreeNode) (*TreeNode, *TreeNode) {

	if node == nil {
		return nil,nil
	}

    leftSubTrees, lp  := cloneTreeNode(node.Left, p)
    rightSubTrees, rp := cloneTreeNode(node.Right, p)

	newNode := TreeNode{Val: node.Val, Left: leftSubTrees, Right: rightSubTrees}

	var trgP *TreeNode

	if lp != nil {
		trgP = lp
	} else if rp != nil {
		trgP = rp
	} else if p != nil && newNode.Val == p.Val {
		trgP = &newNode
	}

	return &newNode, trgP
}

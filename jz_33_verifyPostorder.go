func verifyPostorder(postorder []int) bool {
    
    if len(postorder) <= 1 {
        return true
    }
    
    pIdx := findPartitionIndex(postorder)
    
    for i := pIdx + 1; i < len(postorder)-1; i++ {
        if postorder[i] < postorder[len(postorder) - 1] {
            return false
        }
    }
    
    if pIdx != - 1 {
        if !verifyPostorder(postorder[0:pIdx+1]) {
            return false
        }
    }
    
    if pIdx != (len(postorder) -2) {
        if !verifyPostorder(postorder[pIdx+1:len(postorder) - 1]) {
            return false
        }  
    }
    
    return true
}

func findPartitionIndex(po []int) int {
    
    root := po[len(po) - 1]
    
    pIdx := -1
    for i := 0; i < len(po) - 1; i++ {
        if po[i] < root {
            pIdx = i
        } else {
            break
        }
    }
    return pIdx
}
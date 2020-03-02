func exist(board [][]byte, word string) bool {
    
    if len(board) == 0 {
        return false
    }
    
    t := make([][]bool,len(board))
    for i := 0; i < len(board); i++ {
        t[i] = make([]bool, len(board[i]))
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {   
            
            if board[i][j] != word[0] {
                continue
            }
            
            if len(word) == 1 {
                return true
            }
            
            t[i][j] = true
            if backtracking(board, i, j, t, word[1:]) {
                return true
            } 
            t[i][j] = false
        }
    }
    return false
}

func backtracking(board [][]byte, i, j int, t [][]bool, word string) bool {
    
    curChar := word[0]
    
    if i >= 1 && !t[i-1][j] && board[i-1][j] == curChar {        
        if len(word) == 1 {
            return true
        } else {
            t[i-1][j] = true
            if backtracking(board, i-1, j, t, word[1:]) {
                return true
            }
            t[i-1][j] = false            
        }
    }
    
    if j >= 1 && !t[i][j-1] && board[i][j-1]  == curChar{
        if len(word) == 1 {
            return true
        } else {
            t[i][j-1] = true  
            if backtracking(board, i, j-1, t, word[1:]) {
                return true
            }
            t[i][j-1] = false            
        }
    }
    
    if j < len(board[i]) - 1 && !t[i][j+1] && board[i][j+1] == curChar {
        if len(word) == 1 {
            return true
        } else {
            t[i][j+1] = true
            if backtracking(board,i,j+1, t, word[1:]) {
                return true
            }
            t[i][j+1] = false            
        }
    }
    
    if i < len(board) - 1 && !t[i+1][j] && board[i+1][j] == curChar {
        
        if len(word) == 1 {
            return true
        } else {
            t[i+1][j] = true
            if backtracking(board, i+1, j, t, word[1:]) {
                return true
            }
            t[i+1][j] = false
        }
    }
    
    return false
    
}
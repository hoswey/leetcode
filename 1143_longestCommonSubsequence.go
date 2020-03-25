func longestCommonSubsequence(text1 string, text2 string) int {
    return recurse(text1, text2, len(text1) - 1, len(text2) - 1)
}

func recurse(text1, text2 string, i, j int) int {
    
    if i < 0 || j < 0 {
        return 0
    }
    
    if text1[i] == text2[j] {
        return recurse(text1, text2, i-1, j-1) + 1
    } else {
        return max(recurse(text1, text2, i - 1, j), recurse(text1, text2, i, j-1))
    }    
}

func max(a, b int) int {
    
    if a > b {
        return a
    } else {
        return b
    }
}
	 
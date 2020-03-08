func permutation(s string) []string {
        
    if len(s) == 0 {
        return nil
    }
    
    if len(s) == 1 {
        return []string{s}
    }
    
    var ret []string
    var path []byte
    
    backtrack([]byte(s), &path, &ret)
    
    m := make(map[string]string)
    for _, r := range ret {
        m[r] = r
    }
    
    var ans []string
    for k, _ := range m {
        ans = append(ans, k)
    }

    return ans
}

func backtrack(words []byte, path *[]byte, ans *[]string) {
    
    if len(words) == 1 {        
        dst := make([]byte, len(*path))
        copy(dst, *path)        
        dst = append(dst, words[0])
        // fmt.Printf("ret:%v\n", string(dst))
        *ans = append(*ans, string(dst)) 
        return
    }
    
    for i := 0; i < len(words); i++ {
        
        var remain []byte        
        if i == 0 {
            remain = words[1:len(words)]
        } else if i == len(words) - 1 {
            remain = words[0: len(words) - 1]
        } else {
            remain = append(words[0:i:i], words[i+1:len(words):len(words)]...)
        }
        
        *path = append(*path, words[i])
        // fmt.Printf("Cur: %s, Path:%s, Remain:%s, lenOfWords: %d, i:%d, words:%s\n", string(words[i]), string(*path), string(remain), len(words), i, string(words))
        backtrack(remain, path, ans)
        *path = (*path)[0: len(*path) - 1]
    }    
}
import "container/list"

func isValid(s string) bool {
    
    l := list.New()
    
    for i := 0; i < len(s); i++ {
        
        c := s[i]
        ele := l.Front()
        
        if ele == nil {
            l.PushFront(c)
            continue
        }
        
        pValue := ele.Value.(byte)
        
        if pValue == '(' && c == ')' || pValue == '[' && c == ']' || pValue == '{' && c == '}' {
            l.Remove(ele)
        } else {
            l.PushFront(c)
        }
    }
    
    return l.Len() == 0
}
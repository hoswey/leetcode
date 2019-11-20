//Not dp solution

func isSubsequence(s string, t string) bool {
    
    var ret bool 
    var pos int

    for _, trg := range t {

    	if s[pos] == trg {
    		pos++
    	}

    	if pos == len(s) {
    		ret = true
    		break
    	}
    }

    return ret
}
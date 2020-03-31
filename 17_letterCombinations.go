var letters = make(map[byte][]string)

func init(){
	letters['2'] = []string{"a","b","c"}
	letters['3'] = []string{"d","e","f"}
	letters['4'] = []string{"g","h","i"}
	letters['5'] = []string{"j","k","l"}
	letters['6'] = []string{"m","n","o"}
	letters['7'] = []string{"p","q","r","s"}				
	letters['8'] = []string{"t","u","v"}
	letters['9'] = []string{"w","x","y","z"}	
}

func letterCombinations(digits string) []string {
    
    if digits == "" {
        return nil
    }
    
	var ans []string
	recurse("", []byte(digits), &ans)
	return ans
}

func recurse(prefix string, nextDigits []byte, ans *[]string) {

	if len(nextDigits) == 0 {
		*ans = append(*ans, prefix)
		return
	}

	for _, str := range letters[nextDigits[0]] {
		recurse(prefix + str, nextDigits[1:], ans)
	}
} 
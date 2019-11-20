import (
    "strings"
)

func wordBreak(s string, wordDict []string) bool {

    var arr []byte
    arr = append(arr, ' ')
    arr = append(arr, []byte(s)...)
    
    s = string(arr)

	dp := make([]bool, len(s))	
	dp[0] = true
    
	for i := 1; i < len(s); i++ {

		if dp[i - 1] {
			for _, word := range wordDict {
				if strings.HasPrefix(s[i:], word) {
					dp[i + len(word) - 1] = true
				}
			}
		}
	}

	return dp[len(s) - 1]
}
func generateParenthesis(n int) []string {

	var ans []string
	backtrack(n, n, "", &ans)
	return ans
}

func backtrack(l, r int, path string, ans *[]string) {

	if l > r || l < 0 || r < 0 {
		return
	}

	if l == 0 && r == 0 {
		*ans = append(*ans, path)
		return
	}

	backtrack(l-1, r, path + "(", ans)
	backtrack(l, r-1, path + ")", ans)
}
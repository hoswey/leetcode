func lengthOfLongestSubstring(s string) int {

	var i, j,max int

	m := make(map[byte]byte)
    for i < len(s) && j < len(s) {
		if _, ok := m[s[j]]; !ok {

			m[s[j]] = s[j]
			if j - i + 1 > max {
				max = j - i + 1
			}
			j++
		} else {
			delete(m, s[i])
			i++
		}
	}
    return max
}
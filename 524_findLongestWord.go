/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */
// "abpcplea"
// ["ale","apple","monkey","plea"]
// @lc code=start
import "sort"

func findLongestWord(s string, dictionary []string) string {

	sort.Strings(dictionary)

	var ans string
	for _, w := range dictionary {

      var i, j, cnt int
      for i < len(s) && j < len(w) {
          if s[i] == w[j] {
            i++
            j++
            cnt++
          } else {
            i++
          }
      }

      if cnt == len(w) && len(ans) < len(w) {
        ans = w
      }
	}
	return ans
}
// @lc code=end


/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */

// @lc code=start
import "container/list"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
 
	//build dag
	dag := bfs(beginWord, endWord, wordList)
	var ans [][]string
	var current []string
	current = append(current, beginWord)
	backtracking(dag, beginWord, endWord, current, &ans)

	return ans
}

func backtracking(dag map[string][]string, beginWord string, endWord string, current []string, ans *[][]string) {

	if beginWord == endWord {
		*ans = append(*ans, append([]string(nil), current...))
		return
	}

	nextWords, ok := dag[beginWord]
	if !ok {
		return
	}

	for _, word := range nextWords {
		current = append(current, word)
		backtracking(dag, word, endWord, current, ans)
		current = current[:len(current) - 1]
	}
}

func findNeighbors(word string, wordSet map[string]string) []string {

	var neighbors []string
	for i := 0; i < len(word); i++ {
		for c := 'a'; c <= 'z'; c++ {
			if byte(c) == word[i] {
				continue
			}
			if n, ok := wordSet[word[:i] + string(c) + word[i+1:]]; ok {
				neighbors = append(neighbors, n)
			}
		}
	}
	return neighbors
}

func bfs(beginWord string, endWord string, wordList []string) map[string][]string {

	l := list.New()
	l.PushBack(beginWord)

	wordSet := make(map[string]string)
	for i := 0; i < len(wordList); i++ {
		wordSet[wordList[i]] = wordList[i]
	}
	delete(wordSet, beginWord)

	dag := make(map[string][]string)
	visited := make(map[string]string)
	visited[beginWord] = beginWord

	for l.Len() > 0 {

		var nextTimeVisit []string
		for i := l.Len(); i > 0; i-- {
			e := l.Front()
			l.Remove(e)

			word := e.Value.(string)
			neighbors := findNeighbors(word, wordSet)
			dag[word] = neighbors

			nextTimeVisit = append(nextTimeVisit, neighbors...)
			for _, neighbor := range neighbors {
				if _, ok := visited[neighbor]; !ok {
					visited[neighbor] = neighbor
					l.PushBack(neighbor)
				}
			}
		}

		for _, word := range nextTimeVisit {
			delete(wordSet, word)
		}
	}
	return dag
}
// @lc code=end
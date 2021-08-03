/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	root *Node
}

type Node struct {
	value    bool
	children map[byte]*Node
}

func newNode() *Node {

	return &Node{value: false, children: make(map[byte]*Node)}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: newNode()}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

	children := this.root.children
	for i := 0; i < len(word); i++ {
		c := word[i]
		node, ok := children[c]
		if !ok {
			node = newNode()
			children[c] = node
		}

		if i == len(word)-1 {
			node.value = true
		}
		children = node.children
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	children := this.root.children
	for i := 0; i < len(word); i++ {
		c := word[i]
		node, ok := children[c]
		if !ok {
			return false
		}
		if i == len(word) - 1 {
			return node.value
		}
		children = node.children

	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

	children := this.root.children
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		node, ok := children[c]
		if !ok {
			return false
		}
		if i == len(prefix) - 1 {
			return true
		}
		children = node.children

	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end


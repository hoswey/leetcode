/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

	//org:new
	orgToNew := make(map[*Node]*Node)

	var ans *Node
	cur := head
	for cur != nil {
		node := new(Node)

		if ans == nil {
			ans = node
		}

		orgToNew[cur] = node
		cur = cur.Next
	}

	for k, v := range orgToNew {

		v.Val = k.Val

		v.Next = orgToNew[k.Next]
		v.Random = orgToNew[k.Random]
	}

	return ans
}
// @lc code=end


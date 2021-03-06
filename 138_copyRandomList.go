// 给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

// 要求返回这个链表的 深拷贝。 

// 我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

// val：一个表示 Node.val 的整数。
// random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
 
 
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	cp2Org := make(map[*Node]*Node)
	org2Cp := make(map[*Node]*Node)
	randoms := make(map[*Node]*Node)

	dummmy := new(Node)
    cur := dummmy
	for head != nil {

		cp := new(Node)
		cp.Val = head.Val

		cp2Org[cp] = head
		org2Cp[head] = cp
		randoms[head] = head.Random

        cur.Next = cp
        cur = cur.Next
        head = head.Next
	}

	head = dummmy.Next
	for head != nil {

		org := cp2Org[head]
		random  := randoms[org]
		head.Random = org2Cp[random]

        head = head.Next
	}

	return dummmy.Next
}
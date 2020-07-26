/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 */

// @lc code=start
import "math/rand"

var p = 0.25
var MaxLevel = 32

type Node struct {
	forward []*Node
	val int
}

type Skiplist struct {
	head *Node
	level int
}

func newNode(val int, level int) *Node{
	node := Node{forward: make([]*Node, level), val: val}
	return &node
}

func Constructor() Skiplist {
	return Skiplist{head: newNode(0, MaxLevel), level: 0}
}

func (this *Skiplist) Search(target int) bool {

	x := this.head
	for i := this.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].val < target {
			x = x.forward[i]
		}
	}

	if x.forward[0] != nil && x.forward[0].val == target {
		return true
	}
	return false
}

func (this *Skiplist) Add(num int)  {

	update := make([]*Node, MaxLevel)

	x := this.head
	for i := this.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].val < num {
			x = x.forward[i]
		}
		update[i] = x
	}

	level := randomLevel()
	if level > this.level {
		for i := this.level; i < level; i++ {
			update[i] = this.head
		}
		this.level = level
	}

	added := newNode(num, level)
	for i := 0; i < level; i++ {
		added.forward[i] = update[i].forward[i]
		update[i].forward[i] = added
	}
}

func randomLevel() int {

	level := 1
	for rand.Float64() < p && level < MaxLevel {
		level++
	}

	return level
}


func (this *Skiplist) Erase(num int) bool {

	update := make([]*Node, MaxLevel)
	x := this.head
	for i := this.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].val < num {
			x = x.forward[i]
		}
		update[i] = x
	}

	trg := x.forward[0]
	if trg == nil || trg.val != num {
		return false
	}

	for i := 0; i < this.level; i++ {
		if update[i].forward[i] == nil || update[i].forward[i].val != num {
			break
		}
		update[i].forward[i] = update[i].forward[i].forward[i]
	}

	for this.level > 0 && this.head.forward[this.level - 1] == nil {
		this.level--
	}
	return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end

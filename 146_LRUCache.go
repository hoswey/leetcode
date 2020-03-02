import "container/list"

type LRUCache struct {
	capacity int
	dlist *list.List
	m map[int]*list.Element
	l int
}

type KV struct {
	Key int
	Value int
}

func Constructor(capacity int) LRUCache {

	c := LRUCache{}
	c.capacity = capacity
	c.dlist = list.New()
    c.m = make(map[int]*list.Element)
	return c
}


func (this *LRUCache) Get(key int) int {

	if e, ok := this.m[key]; ok {
		this.dlist.MoveToFront(e)
        return e.Value.(*KV).Value
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {

	if e, ok := this.m[key]; ok {
        kv := e.Value.(*KV)
        kv.Value = value
		this.dlist.MoveToFront(e)
	} else {
		if this.l == this.capacity {
            back := this.dlist.Back()
			this.dlist.Remove(back)
            delete(this.m, back.Value.(*KV).Key)
		} else {
			this.l++
		}
        ne := this.dlist.PushFront(&KV{Key: key, Value: value})
        this.m[key] = ne
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
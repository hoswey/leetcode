import "container/list"

type LFUCache struct {
    
    capacity int
    length int

    entries map[int]*list.Element
    freqs map[int]*list.List

    minFre int
}

type Node struct {

	key int
	value int 
	freq int
}


func Constructor(capacity int) LFUCache {

	cache := LFUCache{}
	cache.capacity = capacity
    cache.entries = make(map[int]*list.Element)
    cache.freqs = make(map[int]*list.List)
	return cache
}


func (this *LFUCache) Get(key int) int {

    if this.capacity == 0 {
        return -1
    }    
    
	if e, ok := this.entries[key]; ok {
		
		node := e.Value.(*Node)

		this.incrFreq(e)
		return node.value
	} else {
		return -1
	}
    
}

func (this *LFUCache) incrFreq(e *list.Element) {

	node := e.Value.(*Node)

	this.freqs[node.freq].Remove(e)
	if this.freqs[node.freq].Len() == 0 {
		delete(this.freqs, node.freq)
		if (this.minFre == node.freq) {
			this.minFre++
		}
	}

	node.freq++
	l, ok := this.freqs[node.freq]

	if !ok {
		l = list.New()
		this.freqs[node.freq] = l
	}
	this.entries[node.key] = l.PushFront(node)
}


func (this *LFUCache) Put(key int, value int)  {

    if this.capacity == 0 {
        return
    }
    
	if e, ok := this.entries[key]; ok {
		node := e.Value.(*Node)
		node.value = value        
 		this.incrFreq(e)
	} else {

		if this.length == this.capacity {
			lfu := this.freqs[this.minFre].Back()
			this.freqs[this.minFre].Remove(lfu)
			if this.freqs[this.minFre].Len() == 0 {
				delete(this.freqs, this.minFre)
			}
			delete(this.entries, lfu.Value.(*Node).key)
		} else {
			this.length++
		}

		l, ok := this.freqs[0]
		if !ok {
			l = list.New()
			this.freqs[0] = l
		}
		this.minFre = 0
		
        inserted := l.PushFront(&Node{key: key, value: value, freq: 0})
		this.entries[key] = inserted
	}
    
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
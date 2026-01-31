package main

import (
	"container/list"
)

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity  int
	evictList *list.List
	cache     map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		evictList: list.New(),
		cache:     make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.evictList.MoveToFront(node)
		// 型別斷言：把 any 轉回 *entry 結構
		return node.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.evictList.MoveToFront(node)
		node.Value.(*entry).value = value
		return
	}

	// 新增節點
	ent := &entry{key, value}
	node := this.evictList.PushFront(ent)
	this.cache[key] = node

	// 超出容量時移除末尾
	if this.evictList.Len() > this.capacity {
		this.removeOldest()
	}
}

func (this *LRUCache) removeOldest() {
	lastNode := this.evictList.Back()
	if lastNode != nil {
		this.evictList.Remove(lastNode)
		kv := lastNode.Value.(*entry)
		delete(this.cache, kv.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

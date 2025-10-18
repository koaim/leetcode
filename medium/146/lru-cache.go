package leetcode

import "container/list"

/*
Implement the LRUCache class:

- LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
- int get(int key) Return the value of the key if the key exists, otherwise return -1.
- void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.
*/

type LRUCache struct {
	data  map[int]*list.Element
	queue *list.List
	cap   int
}

type kv struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{map[int]*list.Element{}, list.New(), capacity}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.data[key]; ok {
		this.queue.MoveToBack(v)
		return v.Value.(kv).value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.data[key]; !ok {
		if len(this.data) < this.cap {
			this.data[key] = this.queue.PushBack(kv{key, value})
		} else {
			e := this.queue.Front()
			this.queue.Remove(e)
			delete(this.data, e.Value.(kv).key)
			this.data[key] = this.queue.PushBack(kv{key, value})
		}
	} else {
		this.queue.MoveToBack(v)
		v.Value = kv{key, value}
	}
}

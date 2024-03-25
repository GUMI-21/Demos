package lrualg

// LRUCache lru算法,淘汰用到最少次数的缓存
type LRUCache struct {
	size       int64
	capacity   int64
	cache      map[int64]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int64
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int64) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int64) LRUCache {
	l := LRUCache{
		cache:    map[int64]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (lru *LRUCache) Get(key int64) int64 {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key int64, value int64) {
	if _, ok := lru.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
			lru.size--
		}
	} else {
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

// 添加节点到链表头
func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// 移除节点
func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 移动节点到链表头部
func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// 移除尾节点
func (lru *LRUCache) removeTail() *DLinkedNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}

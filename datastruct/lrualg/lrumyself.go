package lrualg

type LruSelfStruct struct {
	size       int64
	cap        int64
	cache      map[int64]*LruNode
	head, tail *LruNode
}

type LruNode struct {
	key, value int64
	pre, next  *LruNode
}

func initLruNode(key, value int64) *LruNode {
	return &LruNode{
		key:   key,
		value: value,
	}
}

// GetLruCache 外部获取缓存方法
func GetLruCache(cap int64) *LruSelfStruct {
	lru := &LruSelfStruct{
		size:  0,
		cap:   cap,
		cache: make(map[int64]*LruNode),
		head:  initLruNode(0, 0),
		tail:  initLruNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

// Get 获取缓存值,没有key则返回-1,获取成功时将节点移到head
func (lru *LruSelfStruct) Get(key int64) int64 {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	cache := lru.cache[key]
	lru.moveNodeToHead(cache)
	return cache.value
}

// Put 设置缓存值,key存在时重设值并移到head,不存在时设置新值并校验cap,如果大于容量则删除尾结点
func (lru *LruSelfStruct) Put(key, value int64) {
	if _, ok := lru.cache[key]; !ok {
		node := initLruNode(key, value)
		lru.cache[key] = node
		// 因为这里是新的节点,注意这里要用加到头部!
		lru.addNodeToHead(node)
		lru.size++
		if lru.size > lru.cap {
			lru.delTailNode()
			lru.size--
		}
	} else {
		node := lru.cache[key]
		node.value = value
		lru.moveNodeToHead(node)
	}
}

// 添加节点到头部
func (lru *LruSelfStruct) addNodeToHead(node *LruNode) {
	node.pre = lru.head
	node.next = lru.head.next
	lru.head.next.pre = node
	lru.head.next = node
}

// 将节点移动到头部
func (lru *LruSelfStruct) moveNodeToHead(node *LruNode) {
	lru.delNode(node)
	// 注意这里要用加到头部
	lru.addNodeToHead(node)
}

// 删除节点
func (lru *LruSelfStruct) delNode(node *LruNode) {
	if node.pre != nil {
		node.pre.next = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	}
}

// 删除尾部节点
func (lru *LruSelfStruct) delTailNode() {
	node := lru.tail.pre
	if node.pre != lru.head {
		lru.delNode(node)
		delete(lru.cache, lru.tail.pre.key)
	}
}

package lrualg

import (
	"fmt"
	"testing"
)

func TestGetLruCache(t *testing.T) {
	lru := GetLruCache(5)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	lru.Put(5, 5)
	lru.printLruList()
	lru.Get(4)
	lru.printLruList()
	lru.Put(6, 6)
	lru.printLruList()
}

func (lru *LruSelfStruct) printLruList() {
	node := lru.head.next
	for node != lru.tail {
		fmt.Printf("%d:%d,", node.key, node.value)
		node = node.next

	}
	fmt.Printf("\n")
}

package main

import "fmt"

// LRU Cache实现
// 基于双向链表实现缓存的存储功能，链表本身的插入和移动都能在O(1)中实现
// 链表查找效率较低，引入哈希表，实现快速查找功能

// 链表节点
type LRUCacheNode struct {
	key   int
	value int
	prev  *LRUCacheNode
	next  *LRUCacheNode
}

type LRUCache struct {
	capacity int
	count    int
	head     *LRUCacheNode
	tail     *LRUCacheNode
	m        map[int]*LRUCacheNode // 存放和链表节点的key,用于快速查找对应的链表节点
}

func NewLRUCache(cap int) *LRUCache {
	head := new(LRUCacheNode)
	tail := new(LRUCacheNode)
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: cap,
		count:    0,
		head:     head,
		tail:     tail,
		m:        make(map[int]*LRUCacheNode),
	}
}

// LRU对外提供的查找接口
// 不存在返回-1
// 名字的话把查找到的节点更新至链表头部
func (h *LRUCache) Get(key int) int {
	if h.m[key] == nil {
		return -1
	}
	node := h.m[key]
	h.detachNode(node)
	h.insertToFront(node)
	return node.value
}

// 1、插入的数据不存在,新建节点并检查空间是否足够，把数据插入链表头部
// 2、存在的话更新数据至头部节点
func (h *LRUCache) Put(key, value int) {
	if h.m[key] == nil {
		node := new(LRUCacheNode)
		if h.count == h.capacity {
			h.removeTailLRUNode()
		}
		node.key = key
		node.value = value
		h.m[key] = node
		h.insertToFront(node)
		h.count++ //计数
	} else {
		node := h.m[key]
		h.detachNode(node)
		node.value = value //更新
		h.insertToFront(node)
	}
}

// 分离当前节点
func (h *LRUCache) detachNode(node *LRUCacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 把指定节点插入到链表头
func (h *LRUCache) insertToFront(node *LRUCacheNode) {
	var tmp *LRUCacheNode
	tmp = h.head.next //临时记录第一个节点
	h.head.next = node
	node.prev = h.head

	node.next = tmp
	tmp.prev = node
}

// 删除尾节点
// 当缓存满了之后，删除尾部节点
func (h *LRUCache) removeTailLRUNode() {
	var lastNode *LRUCacheNode // 最后一个节点
	lastNode = h.tail.prev
	h.detachNode(lastNode)
	delete(h.m, lastNode.key)
	h.count--
}

func main() {
	lru := NewLRUCache(3)
	lru.Put(1, 3)
	lru.Put(2, 5)
	lru.Put(3, 7)
	fmt.Println(lru.count)
}

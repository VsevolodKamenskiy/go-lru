package lru

import "sync"

type LRUCache interface {
	Add(key string, value any) bool
	Get(key string) (value any, ok bool)
	Remove(key string) (ok bool)
}

type Node struct {
	key   string
	value interface{}
	prev  *Node
	next  *Node
}

type lru struct {
	capacity int
	cache    map[string]*Node
	head     *Node
	tail     *Node
	mu       sync.RWMutex
}

func NewLRUCache(n int) LRUCache {
	if n < 0 {
		n = 0 // Handle negative capacity
	}
	lru := &lru{
		capacity: n,
		cache:    make(map[string]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (l *lru) Add(key string, value any) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.capacity == 0 {
		return false
	}

	if node, exists := l.cache[key]; exists {
		node.value = value
		l.moveToFront(node)
		return false
	}

	node := &Node{key: key, value: value}
	l.cache[key] = node
	l.addToFront(node)

	if len(l.cache) > l.capacity {
		l.removeLast()
	}

	return true
}

func (l *lru) Get(key string) (any, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	node, exists := l.cache[key]
	if !exists {
		return "", false
	}
	l.moveToFront(node)
	return node.value, true
}

func (l *lru) Remove(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	node, exists := l.cache[key]
	if !exists {
		return false
	}
	l.removeNode(node)
	delete(l.cache, key)
	return true
}

func (l *lru) addToFront(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

func (l *lru) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next, node.prev = nil, nil // Cleanup references
}

func (l *lru) moveToFront(node *Node) {
	if l.head.next == node { // Skip if already at front
		return
	}
	l.removeNode(node)
	l.addToFront(node)
}

func (l *lru) removeLast() {
	last := l.tail.prev
	if last == l.head { // Safety check (empty list)
		return
	}
	l.removeNode(last)
	delete(l.cache, last.key)
}

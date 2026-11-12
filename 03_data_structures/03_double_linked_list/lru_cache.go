package doublelinkedlist

import "errors"

var ErrKeyNotFound = errors.New("key tidak ditemukan di dalam cache")

// entry menyimpan key dan value pasangan cache.
type entry[K comparable, V any] struct {
	key   K
	value V
}

// LRUCache merepresentasikan Least Recently Used Cache menggunakan Doubly Linked List + Hash Map.
// Seluruh operasi Get dan Put berjalan dalam waktu konstan O(1).
type LRUCache[K comparable, V any] struct {
	capacity int
	items    map[K]*Node[entry[K, V]]
	list     *DoublyLinkedList[entry[K, V]]
}

// NewLRUCache membuat cache dengan kapasitas maksimum yang ditentukan.
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*Node[entry[K, V]]),
		list:     NewDoublyLinkedList[entry[K, V]](),
	}
}

// Get mengambil value dari cache. Jika ditemukan, node dipindahkan ke posisi paling depan (most recently used). O(1).
func (c *LRUCache[K, V]) Get(key K) (V, error) {
	node, exists := c.items[key]
	if !exists {
		var zero V
		return zero, ErrKeyNotFound
	}

	c.list.MoveToFront(node)
	return node.Value.value, nil
}

// Put menyimpan atau memperbarui value ke cache. Jika kapasitas penuh, item paling belakang (least recently used) di-evict. O(1).
func (c *LRUCache[K, V]) Put(key K, value V) {
	if node, exists := c.items[key]; exists {
		node.Value.value = value
		c.list.MoveToFront(node)
		return
	}

	if c.list.Size() >= c.capacity {
		// Evict least recently used (tail)
		tailNode := c.list.tail
		if tailNode != nil {
			delete(c.items, tailNode.Value.key)
			c.list.RemoveNode(tailNode)
		}
	}

	newNode := c.list.PushFront(entry[K, V]{key: key, value: value})
	c.items[key] = newNode
}

// Size mengembalikan jumlah item saat ini di cache.
func (c *LRUCache[K, V]) Size() int {
	return c.list.Size()
}

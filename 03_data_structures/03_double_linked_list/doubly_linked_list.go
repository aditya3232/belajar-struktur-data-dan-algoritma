package doublelinkedlist

import (
	"errors"
)

var (
	ErrEmptyList = errors.New("doubly linked list kosong")
)

// Node merepresentasikan elemen dengan dua pointer arah (Prev dan Next).
type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

// DoublyLinkedList adalah linked list dua arah yang memegang pointer head dan tail.
type DoublyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewDoublyLinkedList membuat instans baru Doubly Linked List.
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Size mengembalikan jumlah elemen. O(1).
func (d *DoublyLinkedList[T]) Size() int {
	return d.size
}

// IsEmpty memeriksa apakah list kosong. O(1).
func (d *DoublyLinkedList[T]) IsEmpty() bool {
	return d.size == 0
}

// PushFront menyisipkan elemen di awal list. O(1).
func (d *DoublyLinkedList[T]) PushFront(val T) *Node[T] {
	newNode := &Node[T]{Value: val, Prev: nil, Next: d.head}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.head.Prev = newNode
		d.head = newNode
	}
	d.size++
	return newNode
}

// PushBack menyisipkan elemen di akhir list. O(1).
func (d *DoublyLinkedList[T]) PushBack(val T) *Node[T] {
	newNode := &Node[T]{Value: val, Prev: d.tail, Next: nil}
	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.Next = newNode
		d.tail = newNode
	}
	d.size++
	return newNode
}

// PopFront menghapus elemen paling depan. O(1).
func (d *DoublyLinkedList[T]) PopFront() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, ErrEmptyList
	}

	val := d.head.Value
	d.RemoveNode(d.head)
	return val, nil
}

// PopBack menghapus elemen paling belakang. O(1).
func (d *DoublyLinkedList[T]) PopBack() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, ErrEmptyList
	}

	val := d.tail.Value
	d.RemoveNode(d.tail)
	return val, nil
}

// RemoveNode menghapus node tertentu dalam waktu konstan O(1) karena pointer Prev dan Next sudah tersedia.
func (d *DoublyLinkedList[T]) RemoveNode(node *Node[T]) {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		d.head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		d.tail = node.Prev
	}

	node.Prev = nil
	node.Next = nil
	d.size--
}

// MoveToFront memindahkan node yang sudah ada ke posisi paling depan. O(1).
// Operasi kunci yang digunakan pada algoritma LRU Cache.
func (d *DoublyLinkedList[T]) MoveToFront(node *Node[T]) {
	if node == nil || d.head == node {
		return
	}

	// Lepaskan node dari posisinya saat ini
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		d.tail = node.Prev
	}

	// Pasang di depan head
	node.Prev = nil
	node.Next = d.head
	d.head.Prev = node
	d.head = node
}

// TraverseForward menghasilkan slice dari head ke tail.
func (d *DoublyLinkedList[T]) TraverseForward() []T {
	result := make([]T, 0, d.size)
	curr := d.head
	for curr != nil {
		result = append(result, curr.Value)
		curr = curr.Next
	}
	return result
}

// TraverseBackward menghasilkan slice dari tail ke head.
func (d *DoublyLinkedList[T]) TraverseBackward() []T {
	result := make([]T, 0, d.size)
	curr := d.tail
	for curr != nil {
		result = append(result, curr.Value)
		curr = curr.Prev
	}
	return result
}

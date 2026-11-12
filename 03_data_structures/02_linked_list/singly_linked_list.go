package linkedlist

import (
	"errors"
)

var (
	ErrEmptyList       = errors.New("linked list kosong")
	ErrIndexOutOfRange = errors.New("indeks berada di luar rentang")
)

// Node merepresentasikan satu elemen dalam Singly Linked List.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// SinglyLinkedList adalah struktur data linear dengan pointer satu arah.
type SinglyLinkedList[T comparable] struct {
	head *Node[T]
	size int
}

// NewSinglyLinkedList membuat instans baru Singly Linked List.
func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{head: nil, size: 0}
}

// Size mengembalikan jumlah elemen di dalam list. O(1).
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty memeriksa apakah list dalam keadaan kosong. O(1).
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// InsertHead menyisipkan elemen baru di posisi paling depan (head). O(1).
func (l *SinglyLinkedList[T]) InsertHead(val T) {
	newNode := &Node[T]{Value: val, Next: l.head}
	l.head = newNode
	l.size++
}

// InsertTail menyisipkan elemen baru di posisi paling belakang. O(N).
func (l *SinglyLinkedList[T]) InsertTail(val T) {
	newNode := &Node[T]{Value: val, Next: nil}
	if l.head == nil {
		l.head = newNode
		l.size++
		return
	}

	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	l.size++
}

// InsertAt menyisipkan elemen pada indeks tertentu (0-indexed). O(N).
func (l *SinglyLinkedList[T]) InsertAt(index int, val T) error {
	if index < 0 || index > l.size {
		return ErrIndexOutOfRange
	}
	if index == 0 {
		l.InsertHead(val)
		return nil
	}
	if index == l.size {
		l.InsertTail(val)
		return nil
	}

	curr := l.head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}

	newNode := &Node[T]{Value: val, Next: curr.Next}
	curr.Next = newNode
	l.size++
	return nil
}

// DeleteHead menghapus dan mengembalikan elemen paling depan. O(1).
func (l *SinglyLinkedList[T]) DeleteHead() (T, error) {
	var zero T
	if l.IsEmpty() {
		return zero, ErrEmptyList
	}

	val := l.head.Value
	l.head = l.head.Next
	l.size--
	return val, nil
}

// DeleteTail menghapus dan mengembalikan elemen paling belakang. O(N).
func (l *SinglyLinkedList[T]) DeleteTail() (T, error) {
	var zero T
	if l.IsEmpty() {
		return zero, ErrEmptyList
	}

	if l.head.Next == nil {
		val := l.head.Value
		l.head = nil
		l.size--
		return val, nil
	}

	curr := l.head
	for curr.Next.Next != nil {
		curr = curr.Next
	}

	val := curr.Next.Value
	curr.Next = nil
	l.size--
	return val, nil
}

// DeleteValue menghapus kemunculan pertama elemen dengan nilai target. O(N).
func (l *SinglyLinkedList[T]) DeleteValue(val T) bool {
	if l.IsEmpty() {
		return false
	}

	if l.head.Value == val {
		l.head = l.head.Next
		l.size--
		return true
	}

	curr := l.head
	for curr.Next != nil && curr.Next.Value != val {
		curr = curr.Next
	}

	if curr.Next == nil {
		return false // Nilai tidak ditemukan
	}

	curr.Next = curr.Next.Next
	l.size--
	return true
}

// Search mencari nilai target dan mengembalikan (index, true). O(N).
func (l *SinglyLinkedList[T]) Search(val T) (int, bool) {
	curr := l.head
	idx := 0
	for curr != nil {
		if curr.Value == val {
			return idx, true
		}
		curr = curr.Next
		idx++
	}
	return -1, false
}

// Reverse membalik arah seluruh pointer list secara in-place. O(N) waktu, O(1) ruang.
func (l *SinglyLinkedList[T]) Reverse() {
	var prev *Node[T] = nil
	curr := l.head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

// ToSlice mengubah linked list menjadi slice untuk keperluan inspeksi/testing. O(N).
func (l *SinglyLinkedList[T]) ToSlice() []T {
	result := make([]T, 0, l.size)
	curr := l.head
	for curr != nil {
		result = append(result, curr.Value)
		curr = curr.Next
	}
	return result
}

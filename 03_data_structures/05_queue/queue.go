package queue

import "errors"

var ErrEmptyQueue = errors.New("queue kosong")

// node adalah elemen internal linked-list untuk queue.
type node[T any] struct {
	value T
	next  *node[T]
}

// Queue merepresentasikan struktur data FIFO (First In, First Out)
// berbasis linked-list dengan pointer head dan tail untuk memastikan
// operasi Enqueue dan Dequeue selalu berjalan dalam O(1).
type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// NewQueue membuat antrean baru yang kosong.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{head: nil, tail: nil, size: 0}
}

// Enqueue memasukkan elemen baru ke bagian belakang antrean. O(1).
func (q *Queue[T]) Enqueue(val T) {
	newNode := &node[T]{value: val, next: nil}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

// Dequeue mengambil dan menghapus elemen dari bagian depan antrean. O(1).
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrEmptyQueue
	}

	val := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return val, nil
}

// Front melihat elemen paling depan antrean tanpa menghapusnya. O(1).
func (q *Queue[T]) Front() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrEmptyQueue
	}
	return q.head.value, nil
}

// IsEmpty memeriksa apakah queue kosong. O(1).
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Size mengembalikan jumlah elemen di dalam queue. O(1).
func (q *Queue[T]) Size() int {
	return q.size
}

// Task merepresentasikan pekerjaan dalam task queue.
type Task struct {
	ID      int
	Payload string
}

// SimpleTaskDispatcher mensimulasikan worker queue di mana tugas diproses sesuai urutan kedatangan (FIFO).
func SimpleTaskDispatcher(tasks []Task) []int {
	q := NewQueue[Task]()
	for _, task := range tasks {
		q.Enqueue(task)
	}

	processedOrder := make([]int, 0, len(tasks))
	for !q.IsEmpty() {
		task, _ := q.Dequeue()
		processedOrder = append(processedOrder, task.ID)
	}
	return processedOrder
}

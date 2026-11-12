package circularqueue

import "errors"

var (
	ErrQueueFull       = errors.New("circular queue penuh")
	ErrQueueEmpty      = errors.New("circular queue kosong")
	ErrInvalidCapacity = errors.New("kapasitas harus lebih besar dari 0")
)

// CircularQueue merepresentasikan Ring Buffer dengan kapasitas tetap.
// Memanfaatkan operasi aritmatika modulo (% capacity) untuk membungkus (wrap-around)
// indeks front dan rear tanpa pernah menggeser elemen memori.
type CircularQueue[T any] struct {
	data     []T
	capacity int
	head     int
	tail     int
	size     int
}

// NewCircularQueue menginisialisasi ring buffer dengan ukuran tetap.
func NewCircularQueue[T any](capacity int) (*CircularQueue[T], error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}

	return &CircularQueue[T]{
		data:     make([]T, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
		size:     0,
	}, nil
}

// IsFull memeriksa apakah antrean sudah mencapai kapasitas maksimal. O(1).
func (q *CircularQueue[T]) IsFull() bool {
	return q.size == q.capacity
}

// IsEmpty memeriksa apakah antrean sedang kosong. O(1).
func (q *CircularQueue[T]) IsEmpty() bool {
	return q.size == 0
}

// Size mengembalikan jumlah elemen saat ini. O(1).
func (q *CircularQueue[T]) Size() int {
	return q.size
}

// Capacity mengembalikan kapasitas maksimal buffer. O(1).
func (q *CircularQueue[T]) Capacity() int {
	return q.capacity
}

// Enqueue menambahkan elemen baru ke posisi tail. Menolak jika buffer penuh. O(1).
func (q *CircularQueue[T]) Enqueue(val T) error {
	if q.IsFull() {
		return ErrQueueFull
	}

	q.data[q.tail] = val
	q.tail = (q.tail + 1) % q.capacity
	q.size++
	return nil
}

// EnqueueOverwrite menambahkan elemen baru ke posisi tail. Jika buffer penuh,
// elemen terlama di posisi head akan ditimpa (overwritten), dan head digeser maju. O(1).
// Pola ini umum digunakan pada rolling log dan streaming audio.
func (q *CircularQueue[T]) EnqueueOverwrite(val T) {
	q.data[q.tail] = val
	q.tail = (q.tail + 1) % q.capacity

	if q.IsFull() {
		// Geser head karena elemen terlama ditimpa
		q.head = (q.head + 1) % q.capacity
	} else {
		q.size++
	}
}

// Dequeue mengambil dan menghapus elemen dari posisi head. O(1).
func (q *CircularQueue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrQueueEmpty
	}

	val := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return val, nil
}

// Front melihat elemen tertua (di head) tanpa menghapusnya. O(1).
func (q *CircularQueue[T]) Front() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrQueueEmpty
	}
	return q.data[q.head], nil
}

// Rear melihat elemen terbaru (sebelum tail) tanpa menghapusnya. O(1).
func (q *CircularQueue[T]) Rear() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrQueueEmpty
	}
	idx := (q.tail - 1 + q.capacity) % q.capacity
	return q.data[idx], nil
}

// ToSlice mengembalikan elemen aktif dalam urutan FIFO yang teratur.
func (q *CircularQueue[T]) ToSlice() []T {
	result := make([]T, q.size)
	for i := 0; i < q.size; i++ {
		idx := (q.head + i) % q.capacity
		result[i] = q.data[idx]
	}
	return result
}

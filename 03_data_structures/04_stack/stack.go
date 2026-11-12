package stack

import "errors"

var ErrEmptyStack = errors.New("stack kosong")

// Stack merepresentasikan struktur data LIFO (Last In, First Out) berbasis Slice dinamis.
type Stack[T any] struct {
	elements []T
}

// NewStack membuat instans Stack baru.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elements: make([]T, 0)}
}

// Push memasukkan elemen baru ke puncak tumpukan (Top). O(1) amortized.
func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

// Pop mengambil dan menghapus elemen dari puncak tumpukan. O(1).
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, ErrEmptyStack
	}

	lastIdx := len(s.elements) - 1
	val := s.elements[lastIdx]
	s.elements = s.elements[:lastIdx]
	return val, nil
}

// Peek melihat elemen di puncak tumpukan tanpa menghapusnya. O(1).
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, ErrEmptyStack
	}

	return s.elements[len(s.elements)-1], nil
}

// IsEmpty memeriksa apakah stack dalam keadaan kosong. O(1).
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size mengembalikan jumlah elemen di dalam stack. O(1).
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

// IsValidParentheses adalah aplikasi dunia nyata Stack untuk memvalidasi
// pasangan tanda kurung (), {}, dan [] (digunakan oleh compiler/linter).
func IsValidParentheses(expr string) bool {
	s := NewStack[rune]()
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range expr {
		switch ch {
		case '(', '{', '[':
			s.Push(ch)
		case ')', '}', ']':
			expectedOpen := pairs[ch]
			top, err := s.Pop()
			if err != nil || top != expectedOpen {
				return false
			}
		}
	}

	return s.IsEmpty()
}

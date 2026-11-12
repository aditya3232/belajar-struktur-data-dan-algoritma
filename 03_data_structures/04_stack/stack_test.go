package stack

import (
	"testing"
)

func TestStack_BasicOps(t *testing.T) {
	s := NewStack[int]()

	if !s.IsEmpty() || s.Size() != 0 {
		t.Errorf("Stack baru harus kosong")
	}

	// Pop & Peek pada stack kosong
	_, err := s.Pop()
	if err != ErrEmptyStack {
		t.Errorf("Expected ErrEmptyStack on Pop, got %v", err)
	}

	_, err = s.Peek()
	if err != ErrEmptyStack {
		t.Errorf("Expected ErrEmptyStack on Peek, got %v", err)
	}

	// Push elements
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	top, err := s.Peek()
	if err != nil || top != 30 {
		t.Errorf("Peek expected 30, got %d, err %v", top, err)
	}

	// Pop elements (LIFO order)
	val, err := s.Pop()
	if err != nil || val != 30 {
		t.Errorf("Pop expected 30, got %d", val)
	}

	val, err = s.Pop()
	if err != nil || val != 20 {
		t.Errorf("Pop expected 20, got %d", val)
	}

	val, err = s.Pop()
	if err != nil || val != 10 {
		t.Errorf("Pop expected 10, got %d", val)
	}

	if !s.IsEmpty() {
		t.Errorf("Stack should be empty after popping all elements")
	}
}

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		expected bool
	}{
		{"Valid sederhana", "()", true},
		{"Valid bersarang", "{[()]}", true},
		{"Valid berdampingan", "()[]{}", true},
		{"Valid dengan teks kode", "func main() { arr := []int{1, 2}; }", true},
		{"Invalid kurung tidak seimbang", "(]", false},
		{"Invalid urutan salah", "([)]", false},
		{"Invalid penutup berlebih", "())", false},
		{"Invalid pembuka tidak ditutup", "({[]}", false},
		{"String kosong", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidParentheses(tt.expr)
			if got != tt.expected {
				t.Errorf("IsValidParentheses(%q) = %v, expected %v", tt.expr, got, tt.expected)
			}
		})
	}
}

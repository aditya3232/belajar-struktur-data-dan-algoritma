package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Data acak umum",
			input:    []int{10, 80, 30, 90, 40, 50, 70},
			expected: []int{10, 30, 40, 50, 70, 80, 90},
		},
		{
			name:     "Data sudah terurut",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Data terurut terbalik",
			input:    []int{9, 7, 5, 3, 1},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "Data semua elemen sama",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "Data angka negatif",
			input:    []int{-12, 10, -5, 0, 3, -15},
			expected: []int{-15, -12, -5, 0, 3, 10},
		},
		{
			name:     "Slice kosong",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Satu elemen",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]int, len(tt.input))
			copy(data, tt.input)
			QuickSort(data)
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("QuickSort() got %v, expected %v", data, tt.expected)
			}
		})
	}
}

func TestQuickSort_String(t *testing.T) {
	words := []string{"golang", "python", "rust", "c", "java"}
	QuickSort(words)
	expected := []string{"c", "golang", "java", "python", "rust"}
	if !reflect.DeepEqual(words, expected) {
		t.Errorf("QuickSort(string) got %v, expected %v", words, expected)
	}
}

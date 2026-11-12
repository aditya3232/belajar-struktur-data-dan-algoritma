package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Data acak umum",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "Data sudah terurut",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Data terurut terbalik (worst-case)",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 2}, // intentional check
		},
		{
			name:     "Data dengan duplikasi dan angka negatif",
			input:    []int{3, -2, 5, -2, 0, 1},
			expected: []int{-2, -2, 0, 1, 3, 5},
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

	// Fix test case 3 expected:
	tests[2].expected = []int{1, 2, 3, 4, 5}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]int, len(tt.input))
			copy(data, tt.input)
			BubbleSort(data)
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("BubbleSort() got %v, expected %v", data, tt.expected)
			}
		})
	}
}

func TestBubbleSort_String(t *testing.T) {
	words := []string{"banana", "apple", "cherry", "date"}
	BubbleSort(words)
	expected := []string{"apple", "banana", "cherry", "date"}
	if !reflect.DeepEqual(words, expected) {
		t.Errorf("BubbleSort(string) got %v, expected %v", words, expected)
	}
}

package insertionsort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Data acak umum",
			input:    []int{12, 11, 13, 5, 6},
			expected: []int{5, 6, 11, 12, 13},
		},
		{
			name:     "Data sudah terurut (Best Case O(N))",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Data terurut terbalik (Worst Case)",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Data hampir terurut",
			input:    []int{1, 3, 2, 4, 6, 5},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Data duplikasi dan negatif",
			input:    []int{3, -1, 3, 0, -5, 2},
			expected: []int{-5, -1, 0, 2, 3, 3},
		},
		{
			name:     "Slice kosong",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Satu elemen",
			input:    []int{7},
			expected: []int{7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]int, len(tt.input))
			copy(data, tt.input)
			InsertionSort(data)
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("InsertionSort() got %v, expected %v", data, tt.expected)
			}
		})
	}
}

func TestInsertionSort_String(t *testing.T) {
	input := []string{"zebra", "ant", "monkey", "bear"}
	InsertionSort(input)
	expected := []string{"ant", "bear", "monkey", "zebra"}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("InsertionSort(string) got %v, expected %v", input, expected)
	}
}

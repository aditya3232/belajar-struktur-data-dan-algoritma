package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Data acak umum",
			input:    []int{64, 25, 12, 22, 11},
			expected: []int{11, 12, 22, 25, 64},
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
			name:     "Data duplikasi dan negatif",
			input:    []int{-5, 10, -5, 0, 3, 2},
			expected: []int{-5, -5, 0, 2, 3, 10},
		},
		{
			name:     "Slice kosong",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Satu elemen",
			input:    []int{99},
			expected: []int{99},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]int, len(tt.input))
			copy(data, tt.input)
			SelectionSort(data)
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("SelectionSort() got %v, expected %v", data, tt.expected)
			}
		})
	}
}

func TestSelectionSort_Float(t *testing.T) {
	floats := []float64{3.14, -1.5, 2.718, 0.0, -1.5}
	SelectionSort(floats)
	expected := []float64{-1.5, -1.5, 0.0, 2.718, 3.14}
	if !reflect.DeepEqual(floats, expected) {
		t.Errorf("SelectionSort(float) got %v, expected %v", floats, expected)
	}
}

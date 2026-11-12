package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Data acak umum",
			input:    []int{38, 27, 43, 3, 9, 82, 10},
			expected: []int{3, 9, 10, 27, 38, 43, 82},
		},
		{
			name:     "Data sudah terurut",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Data terurut terbalik",
			input:    []int{9, 8, 7, 6, 5},
			expected: []int{5, 6, 7, 8, 9},
		},
		{
			name:     "Data duplikasi dan negatif",
			input:    []int{-3, 10, -3, 0, 5, -10},
			expected: []int{-10, -3, -3, 0, 5, 10},
		},
		{
			name:     "Slice kosong",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Satu elemen",
			input:    []int{100},
			expected: []int{100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MergeSort() got %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMergeSort_Stability(t *testing.T) {
	// Memastikan stabilitas pada struktur data khusus
	type Item struct {
		key int
		val string
	}

	// Buat fungsi merge khusus item
	mergeItems := func(left, right []Item) []Item {
		res := make([]Item, 0, len(left)+len(right))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if left[i].key <= right[j].key {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}
		res = append(res, left[i:]...)
		res = append(res, right[j:]...)
		return res
	}

	var sortItems func([]Item) []Item
	sortItems = func(items []Item) []Item {
		if len(items) <= 1 {
			return items
		}
		mid := len(items) / 2
		return mergeItems(sortItems(items[:mid]), sortItems(items[mid:]))
	}

	data := []Item{
		{key: 5, val: "first_5"},
		{key: 3, val: "first_3"},
		{key: 5, val: "second_5"},
		{key: 2, val: "first_2"},
	}

	sorted := sortItems(data)
	if sorted[2].val != "first_5" || sorted[3].val != "second_5" {
		t.Errorf("MergeSort must be stable: expected first_5 before second_5")
	}
}

package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	tests := []struct {
		name          string
		target        int
		expectedIndex int
		expectedFound bool
	}{
		{"Target elemen pertama", 11, 0, true},
		{"Target elemen tengah persis", 55, 4, true},
		{"Target elemen terakhir", 99, 8, true},
		{"Target di kuadran kiri", 33, 2, true},
		{"Target di kuadran kanan", 77, 6, true},
		{"Target lebih kecil dari minimum", 5, -1, false},
		{"Target lebih besar dari maksimum", 150, -1, false},
		{"Target di antara dua elemen", 50, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, found := BinarySearch(data, tt.target)
			if found != tt.expectedFound || idx != tt.expectedIndex {
				t.Errorf("BinarySearch(%d) = (%d, %v), expected (%d, %v)",
					tt.target, idx, found, tt.expectedIndex, tt.expectedFound)
			}
		})
	}
}

func TestBinarySearch_EdgeCases(t *testing.T) {
	// Slice kosong
	idx, found := BinarySearch([]int{}, 10)
	if found || idx != -1 {
		t.Errorf("Empty slice search should return (-1, false)")
	}

	// 1 elemen match
	idx, found = BinarySearch([]string{"go"}, "go")
	if !found || idx != 0 {
		t.Errorf("Single element match failed")
	}

	// 1 elemen mismatch
	idx, found = BinarySearch([]string{"go"}, "rust")
	if found || idx != -1 {
		t.Errorf("Single element mismatch failed")
	}
}

package linearsearch

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	data := []int{10, 50, 30, 70, 80, 20}

	tests := []struct {
		name          string
		target        int
		expectedIndex int
		expectedFound bool
	}{
		{"Target di awal", 10, 0, true},
		{"Target di tengah", 30, 2, true},
		{"Target di akhir", 20, 5, true},
		{"Target tidak ditemukan", 999, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, found := LinearSearch(data, tt.target)
			if found != tt.expectedFound || idx != tt.expectedIndex {
				t.Errorf("LinearSearch(%d) = (%d, %v), expected (%d, %v)",
					tt.target, idx, found, tt.expectedIndex, tt.expectedFound)
			}
		})
	}
}

func TestLinearSearch_EdgeCases(t *testing.T) {
	// Slice kosong
	idx, found := LinearSearch([]string{}, "hello")
	if found || idx != -1 {
		t.Errorf("Empty slice should return (-1, false)")
	}

	// 1 elemen
	idx, found = LinearSearch([]string{"golang"}, "golang")
	if !found || idx != 0 {
		t.Errorf("Single element search failed")
	}
}

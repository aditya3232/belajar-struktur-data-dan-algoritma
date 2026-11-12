package interpolationsearch

import (
	"testing"
)

func TestInterpolationSearch_Uniform(t *testing.T) {
	data := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	tests := []struct {
		name          string
		target        int
		expectedIndex int
		expectedFound bool
	}{
		{"Target awal (10)", 10, 0, true},
		{"Target tengah (50)", 50, 4, true},
		{"Target akhir (100)", 100, 9, true},
		{"Target kuadran kanan (80)", 80, 7, true},
		{"Target lebih kecil dari minimum", 5, -1, false},
		{"Target lebih besar dari maksimum", 150, -1, false},
		{"Target di antara nilai tapi tidak ada", 35, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, found := InterpolationSearch(data, tt.target)
			if found != tt.expectedFound || idx != tt.expectedIndex {
				t.Errorf("InterpolationSearch(%d) = (%d, %v), expected (%d, %v)",
					tt.target, idx, found, tt.expectedIndex, tt.expectedFound)
			}
		})
	}
}

func TestInterpolationSearch_EdgeCases(t *testing.T) {
	// Slice kosong
	idx, found := InterpolationSearch([]int{}, 10)
	if found || idx != -1 {
		t.Errorf("Empty slice should return (-1, false)")
	}

	// 1 elemen match
	idx, found = InterpolationSearch([]int{42}, 42)
	if !found || idx != 0 {
		t.Errorf("Single element match failed")
	}

	// 1 elemen mismatch
	idx, found = InterpolationSearch([]int{42}, 99)
	if found || idx != -1 {
		t.Errorf("Single element mismatch failed")
	}

	// Semua elemen bernilai sama
	same := []int{7, 7, 7, 7, 7}
	idx, found = InterpolationSearch(same, 7)
	if !found {
		t.Errorf("All same elements should find target")
	}
	idx, found = InterpolationSearch(same, 10)
	if found || idx != -1 {
		t.Errorf("Should not find missing element in all-same array")
	}
}

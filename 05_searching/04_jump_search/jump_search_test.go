package jumpsearch

import (
	"testing"
)

func TestJumpSearch(t *testing.T) {
	data := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}

	tests := []struct {
		name          string
		target        int
		expectedIndex int
		expectedFound bool
	}{
		{"Target elemen awal (0)", 0, 0, true},
		{"Target tengah (55)", 55, 10, true},
		{"Target elemen akhir (610)", 610, 15, true},
		{"Target di dalam blok (13)", 13, 7, true},
		{"Target tidak ada (di luar rentang atas)", 1000, -1, false},
		{"Target tidak ada (di luar rentang bawah)", -5, -1, false},
		{"Target tidak ada (di antara dua nilai)", 100, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, found := JumpSearch(data, tt.target)
			if found != tt.expectedFound || idx != tt.expectedIndex {
				t.Errorf("JumpSearch(%d) = (%d, %v), expected (%d, %v)",
					tt.target, idx, found, tt.expectedIndex, tt.expectedFound)
			}
		})
	}
}

func TestJumpSearch_EdgeCases(t *testing.T) {
	// Slice kosong
	idx, found := JumpSearch([]int{}, 5)
	if found || idx != -1 {
		t.Errorf("Empty slice should return (-1, false)")
	}

	// 1 elemen match
	idx, found = JumpSearch([]string{"alpha"}, "alpha")
	if !found || idx != 0 {
		t.Errorf("Single element match failed")
	}

	// 1 elemen mismatch
	idx, found = JumpSearch([]string{"alpha"}, "beta")
	if found || idx != -1 {
		t.Errorf("Single element mismatch failed")
	}
}

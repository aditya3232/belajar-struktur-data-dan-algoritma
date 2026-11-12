package exponentialsearch

import (
	"testing"
)

func TestExponentialSearch(t *testing.T) {
	data := []int{2, 3, 4, 10, 40, 50, 60, 70, 80, 90, 100, 120, 140, 160, 180, 200}

	tests := []struct {
		name          string
		target        int
		expectedIndex int
		expectedFound bool
	}{
		{"Target elemen awal persis (2)", 2, 0, true},
		{"Target indeks 1 (3)", 3, 1, true},
		{"Target indeks 4 (40)", 40, 4, true},
		{"Target indeks akhir (200)", 200, 15, true},
		{"Target tengah (90)", 90, 9, true},
		{"Target di luar rentang atas", 999, -1, false},
		{"Target di luar rentang bawah", 1, -1, false},
		{"Target di antara nilai tapi tidak ada", 55, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, found := ExponentialSearch(data, tt.target)
			if found != tt.expectedFound || idx != tt.expectedIndex {
				t.Errorf("ExponentialSearch(%d) = (%d, %v), expected (%d, %v)",
					tt.target, idx, found, tt.expectedIndex, tt.expectedFound)
			}
		})
	}
}

func TestExponentialSearch_EdgeCases(t *testing.T) {
	// Slice kosong
	idx, found := ExponentialSearch([]int{}, 10)
	if found || idx != -1 {
		t.Errorf("Empty slice should return (-1, false)")
	}

	// 1 elemen match
	idx, found = ExponentialSearch([]string{"antigravity"}, "antigravity")
	if !found || idx != 0 {
		t.Errorf("Single element match failed")
	}

	// 1 elemen mismatch
	idx, found = ExponentialSearch([]string{"antigravity"}, "gemini")
	if found || idx != -1 {
		t.Errorf("Single element mismatch failed")
	}
}

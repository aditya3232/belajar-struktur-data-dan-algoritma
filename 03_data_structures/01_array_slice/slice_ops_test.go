package arrayslice

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		index     int
		value     int
		expected  []int
		expectErr bool
	}{
		{
			name:      "Insert di awal",
			initial:   []int{2, 3, 4},
			index:     0,
			value:     1,
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name:      "Insert di tengah",
			initial:   []int{1, 2, 4},
			index:     2,
			value:     3,
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name:      "Insert di akhir",
			initial:   []int{1, 2, 3},
			index:     3,
			value:     4,
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name:      "Insert pada slice kosong",
			initial:   []int{},
			index:     0,
			value:     10,
			expected:  []int{10},
			expectErr: false,
		},
		{
			name:      "Indeks negatif",
			initial:   []int{1, 2},
			index:     -1,
			value:     99,
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "Indeks melebihi batas",
			initial:   []int{1, 2},
			index:     5,
			value:     99,
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InsertAt(tt.initial, tt.index, tt.value)
			if (err != nil) != tt.expectErr {
				t.Fatalf("InsertAt() error = %v, expectErr %v", err, tt.expectErr)
			}
			if !tt.expectErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("InsertAt() got = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestDeleteAt(t *testing.T) {
	tests := []struct {
		name      string
		initial   []string
		index     int
		expected  []string
		expectErr bool
	}{
		{
			name:      "Delete di awal",
			initial:   []string{"a", "b", "c"},
			index:     0,
			expected:  []string{"b", "c"},
			expectErr: false,
		},
		{
			name:      "Delete di tengah",
			initial:   []string{"a", "b", "c"},
			index:     1,
			expected:  []string{"a", "c"},
			expectErr: false,
		},
		{
			name:      "Delete di akhir",
			initial:   []string{"a", "b", "c"},
			index:     2,
			expected:  []string{"a", "b"},
			expectErr: false,
		},
		{
			name:      "Delete slice kosong",
			initial:   []string{},
			index:     0,
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "Indeks di luar batas",
			initial:   []string{"a"},
			index:     1,
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteAt(tt.initial, tt.index)
			if (err != nil) != tt.expectErr {
				t.Fatalf("DeleteAt() error = %v, expectErr %v", err, tt.expectErr)
			}
			if !tt.expectErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("DeleteAt() got = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestSafeSubSlice(t *testing.T) {
	data := []int{10, 20, 30, 40, 50}

	sub, err := SafeSubSlice(data, 1, 4)
	if err != nil {
		t.Fatalf("SafeSubSlice unexpected error: %v", err)
	}

	expected := []int{20, 30, 40}
	if !reflect.DeepEqual(sub, expected) {
		t.Errorf("SafeSubSlice got = %v, expected = %v", sub, expected)
	}

	// Memastikan data terisolasi (deep copy)
	sub[0] = 999
	if data[1] == 999 {
		t.Errorf("SafeSubSlice did not decouple underlying array")
	}

	// Uji error invalid range
	_, err = SafeSubSlice(data, 3, 2)
	if err == nil {
		t.Errorf("Expected error on invalid range start > end")
	}
}

// Benchmark perbandingan alokasi awal (preallocated) vs penambahan dinamis (growing append)
func BenchmarkAppendWithoutPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int
		for j := 0; j < 10000; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAppendWithPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 10000)
		for j := 0; j < 10000; j++ {
			s = append(s, j)
		}
	}
}

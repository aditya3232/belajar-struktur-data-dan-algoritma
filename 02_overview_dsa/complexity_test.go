package overview_dsa

import (
	"testing"
)

// ConstantTime (O(1)): Mengambil elemen pertama array.
// Waktu eksekusi tetap sama berapa pun panjang slice data.
func ConstantTime(data []int) int {
	if len(data) == 0 {
		return 0
	}
	return data[0]
}

// LinearTime (O(N)): Menghitung jumlah seluruh elemen.
// Waktu eksekusi berbanding lurus dengan panjang data.
func LinearTime(data []int) int {
	total := 0
	for _, val := range data {
		total += val
	}
	return total
}

// QuadraticTime (O(N^2)): Menghitung pasangan elemen yang memiliki nilai sama.
// Waktu eksekusi bertumbuh secara kuadratik seiring peningkatan N.
func QuadraticTime(data []int) int {
	duplicatePairs := 0
	n := len(data)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if data[i] == data[j] {
				duplicatePairs++
			}
		}
	}
	return duplicatePairs
}

func TestComplexityFunctions(t *testing.T) {
	sample := []int{1, 2, 3, 2, 1}

	if got := ConstantTime(sample); got != 1 {
		t.Errorf("ConstantTime expected 1, got %d", got)
	}

	if got := LinearTime(sample); got != 9 {
		t.Errorf("LinearTime expected 9, got %d", got)
	}

	// Pasangan duplikat: (index 0 & index 4 -> 1 == 1), (index 1 & index 3 -> 2 == 2) -> 2 pasangan
	if got := QuadraticTime(sample); got != 2 {
		t.Errorf("QuadraticTime expected 2, got %d", got)
	}
}

func BenchmarkConstantTime(b *testing.B) {
	data := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConstantTime(data)
	}
}

func BenchmarkLinearTime(b *testing.B) {
	data := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = LinearTime(data)
	}
}

func BenchmarkQuadraticTime(b *testing.B) {
	// Gunakan data berukuran 1000 agar benchmark selesai dalam batas waktu wajar
	data := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = QuadraticTime(data)
	}
}

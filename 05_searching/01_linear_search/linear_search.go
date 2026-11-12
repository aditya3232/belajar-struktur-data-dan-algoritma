package linearsearch

// LinearSearch mencari nilai target secara sekuensial dari awal hingga akhir slice.
// Tidak memerlukan prasyarat data terurut.
//
// Kompleksitas Waktu:
// - Best Case: O(1) (target berada di indeks pertama)
// - Average Case: O(N)
// - Worst Case: O(N) (target berada di indeks terakhir atau tidak ada)
// Kompleksitas Ruang: O(1)
func LinearSearch[T comparable](arr []T, target T) (int, bool) {
	for i, val := range arr {
		if val == target {
			return i, true
		}
	}
	return -1, false
}

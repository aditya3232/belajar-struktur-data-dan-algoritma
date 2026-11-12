package insertionsort

// Ordered adalah constraint tipe data yang mendukung operator pembanding.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// InsertionSort mengurutkan slice secara in-place dengan cara mengambil satu per satu elemen
// dan menyisipkannya pada posisi yang tepat di dalam sub-array terurut (seperti menyusun kartu di tangan).
//
// Karakteristik kunci: Bersifat adaptif (adaptive). Jika array sudah hampir terurut,
// kompleksitas waktu mendekati O(N).
//
// Kompleksitas Waktu:
// - Best Case: O(N) (data sudah terurut)
// - Average Case: O(N^2)
// - Worst Case: O(N^2) (data terurut terbalik)
// Kompleksitas Ruang: O(1) (In-Place)
// Stabilitas: Stabil (elemen dengan nilai sama mempertahankan urutan relatifnya)
func InsertionSort[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen-elemen sub-array [0..i-1] yang lebih besar dari key ke kanan satu posisi
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

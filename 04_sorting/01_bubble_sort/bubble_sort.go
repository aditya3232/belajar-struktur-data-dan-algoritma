package bubblesort

// Ordered adalah constraint tipe data yang mendukung operator pembanding.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// BubbleSort mengurutkan slice secara in-place dengan membandingkan pasangan elemen bertetangga.
// Menggunakan optimasi flag swapped: jika dalam satu putaran tidak ada elemen yang ditukar,
// algoritma langsung berhenti lebih awal (early termination), menghasilkan waktu O(N) untuk data yang sudah terurut.
//
// Kompleksitas Waktu:
// - Best Case: O(N) (sudah terurut)
// - Average Case: O(N^2)
// - Worst Case: O(N^2) (terurut terbalik)
// Kompleksitas Ruang: O(1) (In-Place)
// Stabilitas: Stabil (elemen duplikat mempertahankan urutan aslinya)
func BubbleSort[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false
		// n - 1 - i karena i elemen terakhir sudah berada di posisi finalnya
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

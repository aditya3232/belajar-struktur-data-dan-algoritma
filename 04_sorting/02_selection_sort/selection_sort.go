package selectionsort

// Ordered adalah constraint tipe data yang mendukung operator pembanding.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// SelectionSort mengurutkan slice secara in-place dengan membagi array menjadi dua bagian:
// bagian terurut (kiri) dan bagian belum terurut (kanan). Pada setiap iterasi, elemen
// terkecil dari bagian belum terurut dipilih dan ditukar ke posisi terdepan bagian belum terurut.
//
// Keunggulan utama: Jumlah operasi penukaran (swap) dibatasi maksimal O(N).
//
// Kompleksitas Waktu:
// - Best Case: O(N^2)
// - Average Case: O(N^2)
// - Worst Case: O(N^2)
// Kompleksitas Ruang: O(1) (In-Place)
// Stabilitas: Tidak Stabil (Unstable secara default)
func SelectionSort[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Hanya tukar jika elemen terkecil bukan di posisi i saat ini
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

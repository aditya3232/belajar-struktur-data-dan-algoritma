package binarysearch

// Ordered adalah constraint tipe data yang mendukung operator pembanding.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// BinarySearch mencari elemen target pada slice yang sudah terurut.
// Membagi area pencarian menjadi dua pada setiap iterasi.
//
// Menggunakan formula titik tengah aman: mid = low + (high - low) / 2
// untuk mencegah potensi bug integer overflow pada array berskala besar.
//
// Prasyarat: Data WAJIB dalam kondisi terurut menaik (sorted ascending).
// Kompleksitas Waktu: O(log N)
// Kompleksitas Ruang: O(1)
func BinarySearch[T Ordered](arr []T, target T) (int, bool) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		// Formula aman pencegah integer overflow
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid, true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1, false
}

package exponentialsearch

// Ordered adalah constraint tipe data yang mendukung operator pembanding.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// ExponentialSearch mencari elemen target pada slice terurut dengan melipatgandakan indeks
// secara eksponensial (1, 2, 4, 8, 16, ...) untuk menemukan rentang batas atas target,
// kemudian mengeksekusi Binary Search pada rentang tersebut.
//
// Karakteristik istimewa:
// Sangat efisien jika elemen target berada dekat dengan awal array, serta dapat diterapkan
// pada array tak terbatas (unbounded / infinite arrays) di mana ukuran akhir data belum diketahui.
//
// Kompleksitas Waktu: O(log i) di mana i adalah indeks posisi target.
// Kompleksitas Ruang: O(1)
func ExponentialSearch[T Ordered](arr []T, target T) (int, bool) {
	n := len(arr)
	if n == 0 {
		return -1, false
	}

	// Cek elemen pertama
	if arr[0] == target {
		return 0, true
	}

	// Tahap 1: Temukan rentang pencarian dengan melipatgandakan i
	i := 1
	for i < n && arr[i] <= target {
		i *= 2
	}

	// Tahap 2: Eksekusi Binary Search pada rentang [i/2 .. min(i, n-1)]
	low := i / 2
	high := min(i, n-1)

	return binarySearch(arr, low, high, target)
}

func binarySearch[T Ordered](arr []T, low, high int, target T) (int, bool) {
	for low <= high {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

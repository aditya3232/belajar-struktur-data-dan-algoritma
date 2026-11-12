package interpolationsearch

// InterpolationSearch mencari target pada array bilangan bulat terurut yang memiliki
// distribusi nilai seragam (uniformly distributed).
//
// Alih-alih selalu membelah tepat di tengah (seperti Binary Search), algoritma ini mengestimasi
// posisi target berdasarkan nilai numeriknya menggunakan rumus interpolasi linear:
//
// pos = low + [ ((target - arr[low]) * (high - low)) / (arr[high] - arr[low]) ]
//
// Kompleksitas Waktu:
// - Best / Average Case: O(log log N) pada data dengan distribusi seragam
// - Worst Case: O(N) jika distribusi data sangat timpang/eksponensial
// Kompleksitas Ruang: O(1)
func InterpolationSearch(arr []int, target int) (int, bool) {
	low := 0
	high := len(arr) - 1

	for low <= high && target >= arr[low] && target <= arr[high] {
		// Jika kedua ujung memiliki nilai yang sama
		if arr[high] == arr[low] {
			if arr[low] == target {
				return low, true
			}
			return -1, false
		}

		// Rumus estimasi posisi interpolasi
		pos := low + int(int64(target-arr[low])*int64(high-low)/int64(arr[high]-arr[low]))

		// Penjagaan batas indeks
		if pos < low || pos > high {
			break
		}

		if arr[pos] == target {
			return pos, true
		} else if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1, false
}

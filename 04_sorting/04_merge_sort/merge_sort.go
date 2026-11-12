package mergesort

// Ordered adalah constraint tipe data yang mendukung operator pembanding.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// MergeSort mengurutkan slice menggunakan paradigma Divide and Conquer.
// Selalu menjamin kompleksitas waktu O(N log N) pada semua kondisi (Best, Average, Worst).
//
// Kompleksitas Waktu: O(N log N)
// Kompleksitas Ruang: O(N) (Out-of-Place karena membutuhkan alokasi memori gabungan)
// Stabilitas: Stabil
func MergeSort[T Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// merge menggabungkan dua sub-slice yang sudah terurut menjadi satu slice terurut baru.
func merge[T Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		// Gunakan <= untuk menjaga stabilitas (elemen kiri diprioritaskan jika nilainya sama)
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Masukkan sisa elemen yang belum habis
	if i < len(left) {
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}

package quicksort

// Ordered adalah constraint tipe data yang mendukung operator pembanding.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// QuickSort mengurutkan slice secara in-place menggunakan algoritma Divide and Conquer
// dengan skema partisi Lomuto yang dioptimasi dengan pemilihan pivot Median of Three
// untuk meminimalkan risiko skenario terburuk O(N^2).
//
// Kompleksitas Waktu:
// - Best Case: O(N log N)
// - Average Case: O(N log N)
// - Worst Case: O(N^2) (jika partisi sangat tidak seimbang)
// Kompleksitas Ruang: O(log N) untuk call stack rekursi
// Stabilitas: Tidak Stabil (Unstable)
func QuickSort[T Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper[T Ordered](arr []T, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

// partition memilih pivot menggunakan Median of Three (low, mid, high)
// dan menata elemen sehingga semua elemen < pivot berada di sebelah kiri,
// dan elemen >= pivot berada di sebelah kanan.
func partition[T Ordered](arr []T, low, high int) int {
	// Optimasi Pivot: Median-of-Three
	mid := low + (high-low)/2
	if arr[mid] < arr[low] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[high] < arr[low] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[high] < arr[mid] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	// Pindahkan median ke posisi high - 1 sebagai pivot
	arr[mid], arr[high] = arr[high], arr[mid]
	pivot := arr[high]

	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

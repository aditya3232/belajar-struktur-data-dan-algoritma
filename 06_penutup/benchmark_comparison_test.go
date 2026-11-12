package penutup

import (
	"math/rand"
	"testing"

	bubblesort "github.com/aditya3232/belajar-struktur-data-dan-algoritma/04_sorting/01_bubble_sort"
	selectionsort "github.com/aditya3232/belajar-struktur-data-dan-algoritma/04_sorting/02_selection_sort"
	insertionsort "github.com/aditya3232/belajar-struktur-data-dan-algoritma/04_sorting/03_insertion_sort"
	mergesort "github.com/aditya3232/belajar-struktur-data-dan-algoritma/04_sorting/04_merge_sort"
	quicksort "github.com/aditya3232/belajar-struktur-data-dan-algoritma/04_sorting/05_quick_sort"

	binarysearch "github.com/aditya3232/belajar-struktur-data-dan-algoritma/05_searching/02_binary_search"
	exponentialsearch "github.com/aditya3232/belajar-struktur-data-dan-algoritma/05_searching/05_exponential_search"
	interpolationsearch "github.com/aditya3232/belajar-struktur-data-dan-algoritma/05_searching/03_interpolation_search"
	jumpsearch "github.com/aditya3232/belajar-struktur-data-dan-algoritma/05_searching/04_jump_search"
	linearsearch "github.com/aditya3232/belajar-struktur-data-dan-algoritma/05_searching/01_linear_search"
)

// GenerateRandomSlice membuat slice data acak untuk pengujian sorting.
func generateRandomSlice(n int) []int {
	rand.Seed(42) // Seed tetap untuk determinisme benchmark
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(n * 10)
	}
	return data
}

// GenerateSortedSlice membuat slice terurut untuk pengujian searching.
func generateSortedSlice(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i * 2 // Distribusi seragam: 0, 2, 4, 6, ...
	}
	return data
}

// ----------------------------------------------------
// BENCHMARK SORTING ALGORITHMS (N = 1000)
// ----------------------------------------------------

func BenchmarkSort_BubbleSort_N1000(b *testing.B) {
	orig := generateRandomSlice(1000)
	data := make([]int, len(orig))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(data, orig)
		bubblesort.BubbleSort(data)
	}
}

func BenchmarkSort_SelectionSort_N1000(b *testing.B) {
	orig := generateRandomSlice(1000)
	data := make([]int, len(orig))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(data, orig)
		selectionsort.SelectionSort(data)
	}
}

func BenchmarkSort_InsertionSort_N1000(b *testing.B) {
	orig := generateRandomSlice(1000)
	data := make([]int, len(orig))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(data, orig)
		insertionsort.InsertionSort(data)
	}
}

func BenchmarkSort_MergeSort_N1000(b *testing.B) {
	orig := generateRandomSlice(1000)
	data := make([]int, len(orig))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(data, orig)
		_ = mergesort.MergeSort(data)
	}
}

func BenchmarkSort_QuickSort_N1000(b *testing.B) {
	orig := generateRandomSlice(1000)
	data := make([]int, len(orig))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(data, orig)
		quicksort.QuickSort(data)
	}
}

// ----------------------------------------------------
// BENCHMARK SEARCHING ALGORITHMS (N = 1000, Target = 1500)
// ----------------------------------------------------

const (
	searchSize   = 1000
	searchTarget = 1500 // berada di indeks 750
)

func BenchmarkSearch_LinearSearch(b *testing.B) {
	data := generateSortedSlice(searchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = linearsearch.LinearSearch(data, searchTarget)
	}
}

func BenchmarkSearch_BinarySearch(b *testing.B) {
	data := generateSortedSlice(searchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = binarysearch.BinarySearch(data, searchTarget)
	}
}

func BenchmarkSearch_InterpolationSearch(b *testing.B) {
	data := generateSortedSlice(searchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = interpolationsearch.InterpolationSearch(data, searchTarget)
	}
}

func BenchmarkSearch_JumpSearch(b *testing.B) {
	data := generateSortedSlice(searchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jumpsearch.JumpSearch(data, searchTarget)
	}
}

func BenchmarkSearch_ExponentialSearch(b *testing.B) {
	data := generateSortedSlice(searchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = exponentialsearch.ExponentialSearch(data, searchTarget)
	}
}

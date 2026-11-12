package jumpsearch

import "math"

// Ordered adalah constraint tipe data yang mendukung operator pembanding.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// JumpSearch mencari nilai target pada slice terurut dengan melompat maju
// sejauh blok optimal m = sqrt(N), kemudian melakukan pencarian linear di dalam blok tersebut.
//
// Prasyarat: Data WAJIB terurut menaik (sorted).
// Keunggulan: Pergerakan utamanya adalah satu arah (forward-only).
//
// Kompleksitas Waktu: O(sqrt(N))
// Kompleksitas Ruang: O(1)
func JumpSearch[T Ordered](arr []T, target T) (int, bool) {
	n := len(arr)
	if n == 0 {
		return -1, false
	}

	// Ukuran lompatan optimal: m = sqrt(n)
	step := int(math.Sqrt(float64(n)))
	if step == 0 {
		step = 1
	}

	prev := 0

	// Melompat ke depan selama elemen di ujung blok masih lebih kecil dari target
	for arr[min(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1, false
		}
	}

	// Pencarian linear di dalam blok [prev .. min(step, n)]
	for arr[prev] < target {
		prev++
		if prev == min(step, n) {
			return -1, false
		}
	}

	if arr[prev] == target {
		return prev, true
	}

	return -1, false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

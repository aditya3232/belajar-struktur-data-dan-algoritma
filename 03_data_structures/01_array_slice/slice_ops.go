package arrayslice

import (
	"errors"
)

var (
	ErrIndexOutOfRange = errors.New("indeks berada di luar rentang")
	ErrEmptySlice      = errors.New("slice kosong")
)

// InsertAt menyisipkan elemen pada indeks tertentu di dalam slice.
// Kompleksitas Waktu: O(N) karena memerlukan pergeseran elemen ke kanan.
func InsertAt[T any](slice []T, index int, value T) ([]T, error) {
	if index < 0 || index > len(slice) {
		return nil, ErrIndexOutOfRange
	}

	// Buat alokasi dengan kapasitas pas (len + 1)
	result := make([]T, len(slice)+1)
	copy(result[:index], slice[:index])
	result[index] = value
	copy(result[index+1:], slice[index:])

	return result, nil
}

// DeleteAt menghapus elemen pada indeks tertentu di dalam slice.
// Kompleksitas Waktu: O(N) karena elemen setelahnya harus digeser ke kiri.
func DeleteAt[T any](slice []T, index int) ([]T, error) {
	if len(slice) == 0 {
		return nil, ErrEmptySlice
	}
	if index < 0 || index >= len(slice) {
		return nil, ErrIndexOutOfRange
	}

	result := make([]T, len(slice)-1)
	copy(result[:index], slice[:index])
	copy(result[index:], slice[index+1:])

	return result, nil
}

// SafeSubSlice memotong slice dan menyalinnya ke buffer baru untuk menghindari
// kebocoran memori (memory leak) akibat memegang referensi ke underlying array besar.
func SafeSubSlice[T any](slice []T, start, end int) ([]T, error) {
	if start < 0 || end > len(slice) || start > end {
		return nil, ErrIndexOutOfRange
	}

	length := end - start
	result := make([]T, length)
	copy(result, slice[start:end])
	return result, nil
}

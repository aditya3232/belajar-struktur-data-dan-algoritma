# 05.2 - Binary Search (Pencarian Biner)

Binary Search adalah algoritma pencarian efisien tingkat tinggi yang bekerja pada koleksi data yang **sudah terurut**. Prinsip utamanya adalah mengeliminasi **setengah** dari area pencarian pada setiap iterasi (*Divide and Conquer*).

```text
Data Terurut: [11, 22, 33, 44, 55, 66, 77, 88, 99], Target: 77

Iterasi 1: Low = 0 (11), High = 8 (99) -> Mid = 4 (55)
           Karena 77 > 55, eliminasi setengah kiri!
           Low = Mid + 1 = 5

Iterasi 2: Low = 5 (66), High = 8 (99) -> Mid = 6 (77)
           arr[6] == 77 (Target Ditemukan! Selesai dalam 2 langkah)
```

Jika array memiliki $1.000.000$ data:
- Linear search membutuhkan hingga **$1.000.000$ perbandingan**.
- Binary search **maksimal hanya membutuhkan $\approx 20$ perbandingan** ($\log_2(1.000.000) \approx 19.93$)!

---

## ⚠️ Waspadai Bug Integer Overflow yang Legendaris

Secara naif, banyak programmer menulis perhitungan titik tengah sebagai:
```go
mid = (low + high) / 2 // BAHAYA!
```
Jika `low + high` melebihi nilai maksimum integer (`math.MaxInt32`), hasil penjumlahan akan meluap (*overflow*) menjadi bilangan negatif dan menyebabkan *panic index out of range*.

Solusi standar industri yang selalu digunakan:
```go
mid = low + (high - low) / 2
```

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

### 1. `git bisect` (Debugging Regresi Kode)
Ketika sistem produksi Anda mendadak mengalami bug setelah ratusan commit baru masuk, menjalankan `git bisect` menggunakan algoritma Binary Search pada commit history untuk menemukan commit pertama yang memicu error hanya dalam hitungan menit.

### 2. Standard Library Go (`slices.BinarySearch`)
Go 1.21+ menyediakan fungsi bawaan `slices.BinarySearch(arr, target)` yang menggunakan algoritma ini untuk pencarian instan pada slice terurut.

### 3. Pencarian Kunci di Node B-Tree Database
Ketika mesin database (seperti MySQL InnoDB) membaca sebuah halaman memori disk (*database page* berukuran 16 KB) yang berisi ratusan index key terurut, database mengeksekusi Binary Search di dalam page tersebut untuk menemukan pointer record berikutnya.

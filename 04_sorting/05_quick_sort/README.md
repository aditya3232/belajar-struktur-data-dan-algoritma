# 04.5 - Quick Sort

Quick Sort adalah algoritma pengurutan paling populer dan paling banyak digunakan di industri. Diciptakan oleh Tony Hoare pada tahun 1959, algoritma ini menggunakan paradigma **Divide and Conquer** dengan memilih sebuah elemen sebagai poros (*pivot*), kemudian mempartisi array ke dalam dua sub-array:
- Elemen yang lebih kecil dari pivot berada di sebelah kiri.
- Elemen yang lebih besar dari pivot berada di sebelah kanan.

```text
Array: [10, 80, 30, 90, 40, 50, 70] (Pivot: 70)
Partisi:
Elemen <= 70: [10, 30, 40, 50] | Pivot: [70] | Elemen > 70: [90, 80]
Rekursi kiri: QuickSort([10, 30, 40, 50])
Rekursi kanan: QuickSort([90, 80])
Hasil Akhir: [10, 30, 40, 50, 70, 80, 90]
```

---

## ⚡ Mengapa Quick Sort Lebih Cepat dari Merge Sort di Praktik?

Secara teori, Merge Sort dan Quick Sort sama-sama memiliki kompleksitas rata-rata $O(N \log N)$. Namun di dunia nyata, Quick Sort hampir selalu lebih cepat 2x hingga 3x karena:
1. **In-Place Memory Access**: Quick Sort tidak memerlukan alokasi buffer array tambahan $O(N)$.
2. **CPU Cache Locality**: Pola akses memorinya sekuensial dan padat, meminimalkan cache miss pada CPU level L1 dan L2.
3. **Konstanta Terkecil**: Instruksi mesin per iterasi pertukaran lebih sedikit.

---

## 🎯 Strategi Pemilihan Pivot (Median-of-Three)

Jika pivot selalu dipilih secara naif (misal elemen terakhir) pada data yang sudah terurut, partisi akan sangat timpang ($1$ vs $N-1$) yang menyebabkan performa jatuh ke **$O(N^2)$**.

Implementasi kami menggunakan **Median-of-Three**:
- Membandingkan elemen pertama (`low`), tengah (`mid`), dan terakhir (`high`).
- Mengambil nilai tengah (*median*) sebagai pivot.
- Menghilangkan kemungkinan skenario terburuk $O(N^2)$ pada data yang sudah terurut.

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

1. **Standar Library Sorting**:
   Fungsi sorting bawaan sistem operasi (fungsi C library `qsort`) dan algoritma sorting modern di Go (`slices.Sort` yang mengadopsi pdqsort) menggunakan Quick Sort sebagai fondasi utamanya.
2. **Game Engine & 3D Computer Graphics**:
   Physics engines menggunakan Quick Sort untuk mengurutkan bounding box poligon objek 3D (Z-sorting / painter's algorithm) untuk rendering grafis real-time 60-120 FPS.

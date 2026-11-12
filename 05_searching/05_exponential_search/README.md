# 05.5 - Exponential Search (Pencarian Eksponensial)

Exponential Search (juga dikenal sebagai *Doubling Search* atau *Galloping Search*) menggabungkan kecepatan eksponensial dalam mempersempit rentang pencarian dengan presisi Binary Search.

---

## 🚀 Cara Kerja Algoritma

Algoritma ini bekerja dalam dua fase:
1. **Fase Penentuan Rentang (Doubling)**:
   - Cek `arr[0]`. Jika cocok, selesai.
   - Mulai dari indeks $i = 1$, lalu lipat gandakan: $i = 1, 2, 4, 8, 16, 32, \dots$ selama $i < N$ dan `arr[i] <= target`.
   - Ketika `arr[i] > target`, kita tahu pasti bahwa target berada di dalam rentang:
     $$[\frac{i}{2} \dots \min(i, N - 1)]$$
2. **Fase Binary Search**:
   - Jalankan Binary Search standar hanya pada rentang kecil $[i/2 \dots \min(i, N-1)]$.

---

## ⚡ Analisis Kompleksitas Waktu

Jika elemen target berada pada indeks ke-$i$:
- Langkah untuk menemukan rentang: $O(\log i)$
- Langkah binary search pada rentang berukuran $i/2$: $O(\log(i/2)) = O(\log i)$
- **Total Kompleksitas Waktu: $O(\log i)$**

> **Poin Kritis**:
> Jika Anda memiliki array berukuran $100.000.000$ elemen, tetapi elemen yang Anda cari berada di indeks ke-10:
> - Binary Search standar membelah array dari indeks $50.000.000$ ($\approx 27$ langkah).
> - Exponential Search hanya butuh $i = 1, 2, 4, 8, 16$ (5 langkah) + binary search di $[8..16]$ (3 langkah) = **hanya 8 langkah**!

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

### 1. Pencarian pada Array Tak Terbatas (Unbounded / Infinite Arrays)
Pada sistem streaming data, data baru terus berdatangan tanpa henti (misal live telemetry logs atau TCP packet streaming buffer) sehingga ukuran total array ($N$) **tidak diketahui / tak berhingga**.
- Binary Search **tidak bisa digunakan** karena membutuhkan batas akhir `high = len(arr) - 1`.
- Exponential Search **dapat digunakan secara sempurna** karena melipatgandakan indeks dari awal sampai menemukan batas atas target tanpa perlu tahu ukuran akhir array!

### 2. Timsort & Galloping Mode
Dalam algoritma pengurutan Timsort (standar di Python dan Java), ketika satu run data secara konsisten selalu menang dalam penggabungan (*merge*), algoritma beralih ke "Galloping Mode" yang menggunakan Exponential Search untuk melompati blok data besar sekaligus.

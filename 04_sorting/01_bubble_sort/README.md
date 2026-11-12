# 04.1 - Bubble Sort

Bubble Sort adalah salah satu algoritma pengurutan berbasis perbandingan yang paling sederhana. Nama "Bubble" berasal dari cara nilai-nilai yang lebih besar "mengapung" secara bertahap ke ujung kanan array seperti gelembung udara di dalam air.

```text
Iterasi 1:
[5, 1, 4, 2, 8] -> Bandingkan 5 & 1 -> Swap: [1, 5, 4, 2, 8]
[1, 5, 4, 2, 8] -> Bandingkan 5 & 4 -> Swap: [1, 4, 5, 2, 8]
[1, 4, 5, 2, 8] -> Bandingkan 5 & 2 -> Swap: [1, 4, 2, 5, 8]
[1, 4, 2, 5, 8] -> Bandingkan 5 & 8 -> Tetap: [1, 4, 2, 5, 8] (8 berada di posisi akhir yang benar)
```

---

## ⚡ Analisis Karakteristik & Kompleksitas

- **Kompleksitas Waktu**:
  - **Best Case**: $O(N)$ (jika array sudah terurut, dengan optimasi flag `swapped`).
  - **Average Case**: $O(N^2)$
  - **Worst Case**: $O(N^2)$ (ketika array terurut terbalik).
- **Kompleksitas Ruang**: $O(1)$ (In-Place, tidak mengalokasikan array baru).
- **Stabilitas (Stability)**: **Stabil**. Elemen dengan nilai yang sama tidak akan bertukar posisi relatif satu sama lain.

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

Meskipun Bubble Sort sangat jarang digunakan dalam sistem skala besar karena $O(N^2)$, algoritma ini memiliki kegunaan nyata:
1. **Edukasi & Fondasi Ilmu Komputer**:
   Memahami konsep perbandingan pasangan, penukaran elemen (*in-place swap*), dan deteksi penghentian awal (*early exit condition*).
2. **Sistem Embedded / Mikrokontroler Keterbatasan Ekstrem**:
   Pada mikrokontroler dengan memori flash instruksi sangat kecil (misal ATtiny dengan flash < 1KB) di mana ukuran data yang disortir hanya 3-6 sensor reading, footprint instruksi Bubble Sort yang sangat ringkas menjadi keuntungan.

# 04.3 - Insertion Sort

Insertion Sort adalah algoritma pengurutan yang bekerja dengan cara membangun deret terurut elemen demi elemen. Konsep kerjanya persis seperti cara manusia mengurutkan kartu remi di tangan: mengambil satu kartu baru dan menyisipkannya di posisi yang tepat di antara kartu-kartu yang sudah terurut.

```text
Awal:        [12, 11, 13, 5, 6]
i=1 (11):    Bandingkan 11 < 12 -> Geser 12 -> [11, 12, 13, 5, 6]
i=2 (13):    13 > 12 -> Tetap di posisi      -> [11, 12, 13, 5, 6]
i=3 (5):     5 lebih kecil dari 13, 12, 11   -> [5, 11, 12, 13, 6]
i=4 (6):     6 disisipkan di antara 5 dan 11 -> [5, 6, 11, 12, 13]
Selesai:     [5, 6, 11, 12, 13]
```

---

## ⚡ Karakteristik & Kompleksitas

- **Kompleksitas Waktu**:
  - **Best Case**: $O(N)$ (jika data sudah terurut).
  - **Average Case**: $O(N^2)$
  - **Worst Case**: $O(N^2)$ (jika data terurut terbalik).
- **Kompleksitas Ruang**: $O(1)$ (In-Place).
- **Adaptif (Adaptive)**: Ya, jumlah operasi berbanding lurus dengan jumlah inversi yang ada.
- **Online Algorithm**: Mampu menyortir sekumpulan data yang diterima secara bertahap (streaming).

---

## 🏢 Mengapa Insertion Sort Sangat Penting di Industri?

Anda mungkin bertanya: *"Mengapa kita mempelajari algoritma $O(N^2)$ jika sudah ada Quick Sort $O(N \log N)$?"*

Jawabannya adalah **Overhead Konstanta Rendah & Cache CPU**:
1. **Sub-rutin pada Algoritma Hybrid Produksi (pdqsort & Timsort)**:
   Di library standar Go (`slices.Sort`), algoritma yang digunakan adalah **pdqsort (Pattern-Defeating Quicksort)**.
   Ketika partisi array rekursif Quick Sort mengecil hingga ukuran $\le 16$ hingga $32$ elemen, fungsi beralih secara otomatis memanggil **Insertion Sort**!
   Sebab pada array kecil, overhead pemanggilan fungsi rekursif Quick Sort jauh lebih lambat daripada perulangan sederhana Insertion Sort yang sangat ramah terhadap cache L1 CPU.
2. **Data Streaming Real-Time**:
   Jika server menerima data yang hampir selalu terurut (misalnya sensor timestamp yang 99% terurut dengan sesekali delay kecil), Insertion Sort berjalan hampir instan dalam $O(N)$.

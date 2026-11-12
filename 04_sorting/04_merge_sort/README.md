# 04.4 - Merge Sort

Merge Sort adalah algoritma pengurutan berbasis strategi **Divide and Conquer** (Bagi dan Taklukkan). Dikembangkan oleh John von Neumann pada tahun 1945, algoritma ini menjadi tolok ukur standar algoritma pengurutan dengan jaminan kompleksitas waktu yang ketat.

```text
                  [38, 27, 43, 3, 9, 82, 10]
                             /        \
              [38, 27, 43, 3]          [9, 82, 10]
                /        \               /      \
            [38, 27]    [43, 3]        [9, 82]   [10]
            /    \      /    \         /    \      |
          [38]  [27]  [43]   [3]     [9]   [82]   [10]
            \    /      \    /         \    /      |
           [27, 38]    [3, 43]        [9, 82]     [10]
               \          /               \        /
              [3, 27, 38, 43]            [9, 10, 82]
                     \                       /
                     [3, 9, 10, 27, 38, 43, 82]
```

---

## ⚡ Analisis Kompleksitas & Karakteristik

- **Kompleksitas Waktu**: **Selalu $O(N \log N)$** di semua kasus (Best Case, Average Case, dan Worst Case). Tidak ada input jahat (*adversarial input*) yang bisa memperlambat Merge Sort menjadi $O(N^2)$.
- **Kompleksitas Ruang**: $O(N)$ (Out-of-Place untuk array karena memerlukan buffer untuk menampung hasil gabungan).
- **Stabilitas (Stability)**: **Stabil**. Sangat aman digunakan untuk tipe data multi-kunci (misal: mengurutkan data pegawai berdasarkan gaji tanpa merusak urutan nama sebelumnya).

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

### 1. External Sorting (Data Skala Raksasa di Database / Hadoop)
Bagaimana cara mengurutkan file log sebesar **500 Gigabyte** jika komputer Anda hanya memiliki **16 Gigabyte RAM**?
- Array tidak bisa dimuat seluruhnya ke RAM (akan langsung memicu OOM).
- Solusi: **K-way Merge Sort**.
  1. Baca potongan data 4 GB dari hard disk ke RAM.
  2. Urutkan potongan di RAM, lalu tulis kembali ke file sementara di disk (*runs*).
  3. Lakukan proses **Merge** bertahap membaca stream dari file-file sementara tersebut langsung ke output disk.

### 2. Pengurutan Linked List
Merge Sort adalah algoritma pengurutan terbaik untuk Linked List karena tidak membutuhkan alokasi memori tambahan $O(N)$ (cukup mengubah pointer `Next`), berbeda dengan Quick Sort yang sulit melakukan partisi acak pada linked list.

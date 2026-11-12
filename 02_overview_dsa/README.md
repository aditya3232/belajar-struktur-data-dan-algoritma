# 02 - Seputar Struktur Data dan Algoritma & Analisis Kompleksitas

Modul ini membahas fondasi teoritis paling krusial dalam ilmu komputer: **bagaimana kita mengukur efisiensi sebuah algoritma dan memilih struktur data yang tepat secara objektif**.

---

## 💡 Apa Itu Struktur Data dan Algoritma?

1. **Struktur Data (Data Structure)**:
   Format khusus untuk mengorganisasi, menyimpan, dan mengelola data di dalam memori komputer agar dapat diakses dan dimodifikasi secara efisien.
   *Contoh*: Array, Linked List, Stack, Queue, Tree, Hash Map.

2. **Algoritma (Algorithm)**:
   Urutan langkah-langkah logis dan terdefinisi dengan jelas untuk menyelesaikan suatu masalah komputasi atau memproses data dari *input* menjadi *output*.
   *Contoh*: Sorting, Searching, Traversal, Dynamic Programming.

> **Analogi**: Struktur data adalah lemari/rak perkakas dengan sekat-sekat tertentu, sedangkan algoritma adalah prosedur atau cara Anda mengambil, menyusun, dan memperbaiki barang menggunakan perkakas tersebut.

---

## ⏱️ Mengapa Tidak Mengukur dengan Detik/Milidetik Saja?

Mengapa kita tidak cukup mengukur kecepatan program dengan fungsi *timer* (seperti `time.Now()` di Go)?
- **Variabilitas Hardware**: Komputer dengan prosesor Intel Core i9/Apple M3 pasti menyelesaikan loop lebih cepat daripada VPS murah 1 vCPU.
- **Beban Sistem**: CPU load, background processes, dan thermal throttling memengaruhi waktu eksekusi.
- **Arsitektur & OS**: Perbedaan arsitektur (ARM64 vs x86_64) dan OS scheduler.

Oleh karena itu, ilmu komputer menggunakan **Analisis Asimtotik (Asymptotic Analysis)** untuk mengukur performa relatif terhadap pertumbuhan ukuran input ($N$), tanpa bergantung pada mesin spesifik.

---

## 📐 Notasi Asimtotik (Asymptotic Notations)

1. **Big-O ($O$) — Upper Bound (Batas Atas)**:
   Menjelaskan skenario terburuk (*worst-case scenario*). Menjamin bahwa algoritma tidak akan berjalan lebih lambat dari batas fungsi ini ketika input membesar. Ini adalah notasi yang paling sering dipakai di industri.
2. **Big-$\Omega$ ($\Omega$) — Lower Bound (Batas Bawah)**:
   Menjelaskan skenario terbaik (*best-case scenario*).
3. **Big-$\Theta$ ($\Theta$) — Tight Bound (Batas Ketat)**:
   Menjelaskan kompleksitas rata-rata (*average case*) ketika batas atas dan batas bawah memiliki laju pertumbuhan yang sama.

---

## 📊 Spektrum Kompleksitas Waktu (Time Complexity)

| Notasi | Nama | Contoh Algoritma / Operasi | Karakteristik Performa |
|---|---|---|---|
| $O(1)$ | Constant Time | Mengakses indeks array `arr[0]`, Push/Pop pada Stack | Instan; durasi tidak terpengaruh jumlah data |
| $O(\log N)$ | Logarithmic Time | Binary Search, BST Search (balanced) | Sangat cepat; membelah input menjadi dua setiap langkah |
| $O(N)$ | Linear Time | Linear Search, iterasi 1 loop sepanjang array | Durasi berbanding lurus dengan jumlah data |
| $O(N \log N)$ | Linearithmic Time | Merge Sort, Quick Sort (average) | Standar emas untuk sorting berbasis perbandingan |
| $O(N^2)$ | Quadratic Time | Nested loops, Bubble Sort, Selection Sort | Mulai lambat secara signifikan untuk $N > 10.000$ |
| $O(2^N)$ | Exponential Time | Rekursi Fibonacci naif, brute-force subset | Eksplosif; tidak layak untuk input besar ($N > 30$) |
| $O(N!)$ | Factorial Time | Traveling Salesperson Problem (brute-force) | Mustahil untuk komputasi praktis ($N > 12$) |

---

## 💾 Kompleksitas Ruang (Space Complexity)

Kompleksitas ruang mengukur jumlah memori tambahan (*auxiliary space*) yang dibutuhkan algoritma seiring pertumbuhan $N$:
- **In-Place ($O(1)$ Space)**: Algoritma hanya meminjam beberapa variabel sementara tanpa menyalin array (misal: Bubble Sort, Binary Search iteratif).
- **Out-of-Place ($O(N)$ Space)**: Algoritma mengalokasikan slice atau struktur baru seukuran input (misal: Merge Sort).

---

## 🏢 Penerapan di Dunia Nyata (Real-World Impact)

### 1. Masalah Query N+1 dan Nested Loop $O(N^2)$ di Microservices
Bayangkan Anda memiliki $10.000$ transaksi pengguna.
- Jika Anda mencari data merchant untuk setiap transaksi menggunakan linear search di dalam nested loop ($N \times M$), sistem akan melakukan $10.000 \times 10.000 = 100.000.000$ operasi. Ini bisa memakan waktu beberapa detik dan membuat CPU VPS 100%.
- Mengubah pencarian menjadi Hash Map / Pre-sorted Binary Search menurunkan komputasi menjadi $O(N + M)$ atau $O(N \log M)$, selesai dalam hitungan milidetik.

### 2. Penghematan Biaya Cloud Server (AWS / GCP / Cloudflare)
Di arsitektur *serverless* (AWS Lambda / Google Cloud Run), Anda ditagih per milidetik eksekusi dan memori yang dialokasikan. Algoritma $O(N \log N)$ vs $O(N^2)$ secara langsung membedakan antara tagihan \$50/bulan vs \$5.000/bulan pada volume data harian yang tinggi.

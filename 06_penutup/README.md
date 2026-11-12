# 06 - Penutup: Sintesis, Cheat Sheet & Panduan Arsitektur DSA

Selamat! Anda telah menyelesaikan seluruh kurikulum pembelajaran fundamental **Struktur Data dan Algoritma (DSA) menggunakan bahasa Go (Golang)**.

Modul penutup ini menyajikan sintesis holistik seluruh materi yang telah dipelajari, matriks perbandingan kompleksitas (*Cheat Sheet*), serta panduan praktis pengambilan keputusan arsitektur sistem di dunia industri.

---

## 📊 Matriks Ringkasan Kompleksitas (Cheat Sheet)

### 1. Struktur Data

| Struktur Data | Akses Indeks | Pencarian (Search) | Penyisipan (Insert) | Penghapusan (Delete) | Kompleksitas Ruang | Kasus Terbaik Penggunaan |
|---|---|---|---|---|---|---|
| **Array / Slice** | $O(1)$ | $O(N)$ | $O(N)$ (akhir: $O(1)$ amortized) | $O(N)$ | $O(N)$ | Pembacaan acak sering, buffer I/O, cache-friendly CPU |
| **Singly Linked List** | $O(N)$ | $O(N)$ | $O(1)$ (di head) | $O(1)$ (di head) | $O(N)$ | Sering insert/delete di depan, separate chaining hash table |
| **Doubly Linked List** | $O(N)$ | $O(N)$ | $O(1)$ (head/tail) | $O(1)$ (jika pegang node) | $O(N)$ | LRU Cache, navigasi bolak-balik (Undo/Redo, browser) |
| **Stack (LIFO)** | - | $O(N)$ | $O(1)$ (Push) | $O(1)$ (Pop) | $O(N)$ | Call stack, validasi kurung, DFS, evaluasi ekspresi |
| **Queue (FIFO)** | - | $O(N)$ | $O(1)$ (Enqueue) | $O(1)$ (Dequeue) | $O(N)$ | Task dispatcher, BFS, penanganan antrean request |
| **Circular Queue** | $O(1)$ | $O(N)$ | $O(1)$ | $O(1)$ | $O(K)$ (Fixed) | Streaming audio/video, kernel network ring buffer, zero-allocation |
| **Binary Search Tree** | - | $O(\log N)$ | $O(\log N)$ | $O(\log N)$ | $O(N)$ | Data hierarkis, database indexing, range query |

---

### 2. Algoritma Pengurutan (Sorting)

| Algoritma | Waktu Terbaik (Best) | Waktu Rata-rata (Avg) | Waktu Terburuk (Worst) | Ruang (Space) | Stabil? | Penerapan Utama di Industri |
|---|---|---|---|---|---|---|
| **Bubble Sort** | $O(N)$ | $O(N^2)$ | $O(N^2)$ | $O(1)$ | Ya | Edukasi konsep swap, mikrokontroler flash ultra-sempit |
| **Selection Sort** | $O(N^2)$ | $O(N^2)$ | $O(N^2)$ | $O(1)$ | Tidak | Minimasi jumlah operasi penulisan memori flash ($N$ writes) |
| **Insertion Sort** | $O(N)$ | $O(N^2)$ | $O(N^2)$ | $O(1)$ | Ya | Data hampir terurut, sub-rutin pdqsort / Timsort ($N \le 32$) |
| **Merge Sort** | $O(N \log N)$ | $O(N \log N)$ | $O(N \log N)$ | $O(N)$ | Ya | External sorting (file gigabyte/terabyte di disk), sorting linked list |
| **Quick Sort** | $O(N \log N)$ | $O(N \log N)$ | $O(N^2)$ | $O(\log N)$ | Tidak | Algoritma sorting bawaan sistem (`slices.Sort`), game engine |

---

### 3. Algoritma Pencarian (Searching)

| Algoritma | Prasyarat Data | Kompleksitas Waktu | Kompleksitas Ruang | Skenario Keunggulan Utama |
|---|---|---|---|---|
| **Linear Search** | Bebas (tidak perlu terurut) | $O(N)$ | $O(1)$ | Array kecil ($N < 30$), data acak tidak terstruktur |
| **Binary Search** | Wajib Terurut | $O(\log N)$ | $O(1)$ | Standar emas data terurut, database index lookup, `git bisect` |
| **Interpolation Search** | Wajib Terurut & Distribusi Seragam | $O(\log \log N)$ | $O(1)$ | Data sensor/telemetri numerik dengan kenaikan linear seragam |
| **Jump Search** | Wajib Terurut | $O(\sqrt{N})$ | $O(1)$ | Media penyimpanan dengan operasi forward-only seek (tape/stream) |
| **Exponential Search** | Wajib Terurut | $O(\log i)$ | $O(1)$ | Array tak terbatas (unbounded stream), target dekat dengan awal array |

---

## 🧭 Panduan Keputusan Arsitektur (Architectural Decision Guide)

Saat merancang arsitektur perangkat lunak di dunia nyata, gunakan diagram panduan keputusan berikut:

```text
Kebutuhan Anda:
│
├── Menyimpan & Mengakses Data?
│   ├── Butuh akses indeks acak secepat kilat? ────────────► Slice / Array
│   ├── Kapasitas tetap & anti-alokasi RAM berulang? ──────► Circular Queue (Ring Buffer)
│   ├── Sering insert/delete di awal atau tengah? ─────────► Doubly Linked List
│   ├── Butuh urutan pemrosesan LIFO? ─────────────────────► Stack
│   ├── Butuh urutan pemrosesan FIFO? ─────────────────────► Queue
│   └── Butuh hierarki atau range search terurut? ─────────► Tree (BST / B-Tree)
│
├── Mengurutkan Data?
│   ├── Ukuran sangat besar dan tidak muat di RAM? ────────► Merge Sort (External Sorting)
│   ├── Array umum di dalam memori RAM? ───────────────────► Quick Sort (atau pdqsort)
│   └── Potongan data sangat kecil (N <= 24)? ─────────────► Insertion Sort
│
└── Mencari Data?
    ├── Data tidak terurut? ───────────────────────────────► Linear Search
    ├── Data terurut dan terdistribusi seragam? ───────────► Interpolation Search
    ├── Data streaming tak terbatas (unbounded)? ──────────► Exponential Search
    └── Data terurut umum di memory? ──────────────────────► Binary Search
```

---

## 🚀 Rekomendasi Topik Pembelajaran Lanjutan

Setelah menguasai fondasi dalam repositori ini, langkah strategis berikutnya untuk memperdalam ilmu sistem dan rekayasa perangkat lunak adalah:
1. **Graph Algorithms**: BFS, DFS, Dijkstra Shortest Path, Topological Sort, A* Search.
2. **Hash Tables & Hash Maps**: MurmurHash, FNV-1a, Swiss Tables (SIMD parallel probing), Robin Hood hashing.
3. **Advanced Trees**: Self-Balancing Trees (AVL Tree, Red-Black Tree), B-Tree / B+ Tree, Trie (Prefix Tree), Segment Tree.
4. **Priority Queue & Heap**: Min-Heap, Max-Heap, Fibonacci Heap (digunakan pada task scheduler dan Dijkstra).
5. **Dynamic Programming & Greedy**: Memoization, Tabulation, Knapsack Problem, Interval Scheduling.

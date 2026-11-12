# 🚀 Belajar Struktur Data dan Algoritma (Golang)

Repositori ini berisi kurikulum pembelajaran komprehensif mengenai **Struktur Data dan Algoritma (DSA)** yang diimplementasikan menggunakan bahasa pemrograman **Go (Golang)** modern (Go 1.20+). 

Setiap topik dirancang dengan standar industri:
- Implementasi murni dan idiomatik Go (memanfaatkan *Generics* dan alokasi memori yang efisien).
- Disertai **Unit Test** lengkap (`*_test.go`) mencakup skenario umum dan kasus ekstrem (*edge cases*).
- Disertai penjelasan mendalam tentang **Penerapan di Dunia Nyata (*Real-World Applications*)** di level sistem backend, database, sistem operasi, dan cloud computing.

---

## 📂 Daftar Modul & Navigasi

### 1. Fondasi & Pengenalan
- [01 - Pengantar & Filosofi DSA di Go](./01_intro/README.md)
- [02 - Analisis Kompleksitas Asimtotik (Big-O, Big-$\Omega$, Big-$\Theta$)](./02_overview_dsa/README.md)

### 2. Struktur Data (Data Structures)
- [03.1 - Array dan Slice](./03_data_structures/01_array_slice/README.md) (Cache Locality, Slice Header, Pencegahan Memory Leak)
- [03.2 - Singly Linked List](./03_data_structures/02_linked_list/README.md) (Insert, Delete, Search, Reverse)
- [03.3 - Doubly Linked List & LRU Cache](./03_data_structures/03_double_linked_list/README.md) (Navigasi Dua Arah, $O(1)$ Node Removal, LRU Cache)
- [03.4 - Stack (LIFO)](./03_data_structures/04_stack/README.md) (Push, Pop, Peek, Parentheses Validator)
- [03.5 - Queue (FIFO)](./03_data_structures/05_queue/README.md) (Enqueue, Dequeue, Task Dispatcher)
- [03.6 - Circular Queue (Ring Buffer)](./03_data_structures/06_circular_queue/README.md) (Modulo Indexing, Overwrite Buffer, Zero-Allocation)
- [03.7 - Tree (Binary Search Tree - BST)](./03_data_structures/07_tree/README.md) (Insert, Search, Delete, In-Order, Pre-Order, Post-Order, BFS Level-Order)

### 3. Algoritma Pengurutan (Sorting)
- [04.1 - Bubble Sort](./04_sorting/01_bubble_sort/README.md) (Perbandingan Tetangga & Early Exit Optimization)
- [04.2 - Selection Sort](./04_sorting/02_selection_sort/README.md) (Minimasi Siklus Write $O(N)$ untuk EEPROM/Flash Memory)
- [04.3 - Insertion Sort](./04_sorting/03_insertion_sort/README.md) (Sorting Adaptif & Sub-rutin pdqsort)
- [04.4 - Merge Sort](./04_sorting/04_merge_sort/README.md) (Divide & Conquer, Kestabilan, External Sorting untuk Data Raksasa)
- [04.5 - Quick Sort](./04_sorting/05_quick_sort/README.md) (Median-of-Three Pivot, Lomuto Partition, Cache Friendly)

### 4. Algoritma Pencarian (Searching)
- [05.1 - Linear Search](./05_searching/01_linear_search/README.md) (Pencarian Sekuensial untuk Data Tak Terurut / Ukuran Kecil)
- [05.2 - Binary Search](./05_searching/02_binary_search/README.md) (Eliminasi Setengah Data $O(\log N)$, Safe Midpoint Overflow Bug Prevention)
- [05.3 - Interpolation Search](./05_searching/03_interpolation_search/README.md) (Estimasi Posisi Linear $O(\log \log N)$ untuk Data Sensor/Telemetri Seragam)
- [05.4 - Jump Search](./05_searching/04_jump_search/README.md) (Lompatan Blok Optimal $m = \sqrt{N}$, Forward-Only Seeking)
- [05.5 - Exponential Search](./05_searching/05_exponential_search/README.md) (Pencarian Eksponensial $O(\log i)$ untuk Array Tak Terbatas / Live Stream)

### 5. Evaluasi & Sintesis
- [06 - Penutup, Matriks Kompleksitas & Panduan Keputusan Arsitektur](./06_penutup/README.md)

---

## 🧪 Menjalankan Pengujian (Testing)

### Menjalankan Seluruh Unit Test:
```bash
go test -v ./...
```

### Menjalankan Test pada Modul Spesifik:
```bash
go test -v ./03_data_structures/03_double_linked_list/...
```

### Menjalankan Benchmark Komparatif Seluruh Algoritma:
```bash
go test -bench=. -benchmem ./06_penutup/...
```

---

## 📌 Status Pelacakan Tugas (Issue Tracker)
Seluruh 20 issue roadmap telah diselesaikan dan dapat dilihat pada file [`issue.md`](./issue.md) serta pada [GitHub Issues Repository](https://github.com/aditya3232/belajar-struktur-data-dan-algoritma/issues?q=is%3Aissue).

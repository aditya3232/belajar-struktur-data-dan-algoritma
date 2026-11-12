# Roadmap & Task Tracker: Belajar Struktur Data dan Algoritma (Golang)

Dokumen ini berfungsi sebagai pelacak tugas (*task tracker/issue list*) untuk pembuatan materi pembelajaran Struktur Data dan Algoritma (DSA) menggunakan bahasa pemrograman **Go (Golang)**. Setiap modul dilengkapi dengan teori, implementasi kode idiomatik Go, **unit testing** (`testing` package), serta **studi kasus & penerapan di dunia nyata**.

---

## 🎯 Target & Standar Kualitas

- **Bahasa**: Go (versi modern 1.20+)
- **Testing**: Setiap struktur data dan algoritma wajib memiliki file `*_test.go` dengan cakupan skenario normal, skenario ekstrem (empty, 1 elemen, duplicate, negative), serta benchmark (`testing.B`) jika relevan.
- **Penerapan Dunia Nyata**: Penjelasan mendalam mengapa struktur data / algoritma tersebut relevan di industri (sistem backend, database, OS, network, dsb).
- **Struktur Direktori Proyek**:
  ```text
  .
  ├── go.mod
  ├── issue.md
  ├── 01_intro/
  ├── 02_overview_dsa/
  ├── 03_data_structures/
  │   ├── 01_array_slice/
  │   ├── 02_linked_list/
  │   ├── 03_double_linked_list/
  │   ├── 04_stack/
  │   ├── 05_queue/
  │   ├── 06_circular_queue/
  │   └── 07_tree/
  ├── 04_sorting/
  │   ├── 01_bubble_sort/
  │   ├── 02_selection_sort/
  │   ├── 03_insertion_sort/
  │   ├── 04_merge_sort/
  │   └── 05_quick_sort/
  ├── 05_searching/
  │   ├── 01_linear_search/
  │   ├── 02_binary_search/
  │   ├── 03_interpolation_search/
  │   ├── 04_jump_search/
  │   └── 05_exponential_search/
  └── 06_penutup/
  ```

---

## 📋 Daftar Issues / Modul Pembelajaran

### Milestone 1: Fondasi & Pengenalan (Intro & Overview)

- [x] **ISSUE-01: Inisialisasi Proyek & Intro Pembelajaran**
  - **Deskripsi**: Setup awal repository, konfigurasi `go.mod`, panduan eksekusi test, dan roadmap belajar.
  - **Materi**:
    - Filosofi Go dalam DSA (pointer, memory management, zero value, interface/generics).
    - Cara menjalankan kode dan unit test (`go test -v ./...`, `go test -bench=.`).
  - **Acceptance Criteria**:
    - [x] File `go.mod` terinisialisasi.
    - [x] `01_intro/README.md` menjelaskan roadmap dan cara setup.
  - **Penerapan Dunia Nyata**: Standardisasi repository project di level industri (CI/CD pipeline test readiness).

---

- [x] **ISSUE-02: Seputar Struktur Data dan Algoritma & Analisis Kompleksitas**
  - **Deskripsi**: Fondasi teoritis tentang efisiensi komputasi.
  - **Materi**:
    - Mengapa engineer butuh DSA (trade-off memori vs kecepatan).
    - Asymptotic Notation: Big-O, Big-$\Omega$, Big-$\Theta$.
    - Time Complexity vs Space Complexity.
    - Analisis kasus: Best Case, Average Case, Worst Case.
  - **Acceptance Criteria**:
    - [x] Dokumen penjelasan komprehensif di `02_overview_dsa/README.md`.
    - [x] Contoh kode pembanding kompleksitas $O(1)$, $O(N)$, $O(N^2)$ beserta benchmark.
  - **Penerapan Dunia Nyata**:
    - Menghindari bottleneck performa di microservices ber-traffic tinggi.
    - Estimasi dan pemangkasan biaya cloud server (RAM & CPU usage optimization).

---

### Milestone 2: Struktur Data Linear & Non-Linear

- [x] **ISSUE-03: Array & Slice di Go**
  - **Deskripsi**: Memahami representasi memori kontinu dan perbedaannya di Go.
  - **Materi**:
    - Fixed-size Array vs Dynamic Slice (Slice Header: Pointer, Length, Capacity).
    - Mekanisme pertumbuhan kapasitas (`append`, capacity doubling, allocation overhead).
    - Cache Locality dan pengaruhnya pada performa CPU L1/L2 cache.
  - **Deliverable**:
    - [x] `03_data_structures/01_array_slice/slice_ops.go`
    - [x] `03_data_structures/01_array_slice/slice_ops_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Buffer I/O (`bytes.Buffer`, streaming socket chunk).
    - Pemrosesan matriks gambar/pixel (RGBA manipulation).
    - Batch database record fetcher.

---

- [x] **ISSUE-04: Singly Linked List**
  - **Deskripsi**: Implementasi rantai node satu arah dengan alokasi memori dinamis.
  - **Materi**:
    - Struktur `Node` (Value, Next pointer).
    - Operasi dasar: `InsertHead`, `InsertTail`, `InsertAt`, `DeleteValue`, `Search`, `Reverse`.
    - Kelebihan/kekurangan dibanding Array/Slice.
  - **Deliverable**:
    - [x] `03_data_structures/02_linked_list/singly_linked_list.go`
    - [x] `03_data_structures/02_linked_list/singly_linked_list_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Resolusi tabrakan pada Hash Map (Separate Chaining).
    - Sistem alokasi memori (*free memory block tracking* pada embedded OS).
    - Operasi Undo berantai sederhana.

---

- [x] **ISSUE-05: Double Linked List (Doubly Linked List)**
  - **Deskripsi**: Rantai node dua arah dengan pointer maju dan mundur.
  - **Materi**:
    - Struktur `Node` (Value, Prev, Next).
    - Operasi: `PushFront`, `PushBack`, `RemoveNode`, `TraverseForward`, `TraverseBackward`.
    - Penghapusan $O(1)$ jika referensi node sudah dipegang.
  - **Deliverable**:
    - [x] `03_data_structures/03_double_linked_list/doubly_linked_list.go`
    - [x] `03_data_structures/03_double_linked_list/doubly_linked_list_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Algoritma **LRU Cache (Least Recently Used)**: dikombinasikan dengan Hash Map (`container/list` di Go standard library).
    - Navigasi Back/Forward pada web browser.
    - Playlist music player (Next track / Previous track).

---

- [x] **ISSUE-06: Stack (LIFO - Last In, First Out)**
  - **Deskripsi**: Koleksi elemen dengan mekanisme tumpukan.
  - **Materi**:
    - Operasi: `Push`, `Pop`, `Peek/Top`, `IsEmpty`, `Size`.
    - Dua varian implementasi: Slice-based vs Linked List-based.
  - **Deliverable**:
    - [x] `03_data_structures/04_stack/stack.go`
    - [x] `03_data_structures/04_stack/stack_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Call Stack eksekusi fungsi di Virtual Machine / Runtime Go (goroutine stack).
    - Pengecekan sintaks validitas tanda kurung (Parentheses Validator di linter/compiler).
    - Fitur Undo/Redo di text editor (VS Code, Word).
    - Algoritma Depth-First Search (DFS) & evaluasi ekspresi matematika (Post-fix/Reverse Polish Notation).

---

- [x] **ISSUE-07: Queue (FIFO - First In, First Out)**
  - **Deskripsi**: Koleksi elemen dengan mekanisme antrean masuk di belakang dan keluar di depan.
  - **Materi**:
    - Operasi: `Enqueue`, `Dequeue`, `Front`, `IsEmpty`, `Size`.
    - Masalah pergeseran memori (*shifting cost* $O(N)$ jika menggunakan slice naif vs $O(1)$ linked list).
  - **Deliverable**:
    - [x] `03_data_structures/05_queue/queue.go`
    - [x] `03_data_structures/05_queue/queue_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Antrean pesan / Task Queue asynchronous (asumsi RabbitMQ, Kafka, AWS SQS consumer).
    - Web Server Connection pool & request handling buffer.
    - Breadth-First Search (BFS) pada graf atau pohon.
    - Sistem antrean cetak (Print spooler) dan OS CPU scheduling (Round Robin ready queue).

---

- [x] **ISSUE-08: Circular Queue (Ring Buffer)**
  - **Deskripsi**: Antrean dengan ukuran tetap (*fixed capacity*) di mana posisi akhir terhubung kembali ke posisi awal.
  - **Materi**:
    - Front & Rear index manipulation menggunakan operasi modulo (`% capacity`).
    - Deteksi kondisi `IsFull` dan `IsEmpty`.
    - Zero-allocation setelah inisialisasi awal.
  - **Deliverable**:
    - [x] `03_data_structures/06_circular_queue/circular_queue.go`
    - [x] `03_data_structures/06_circular_queue/circular_queue_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Buffer streaming audio/video (mencegah stuttering tanpa realokasi memori berulang).
    - Network card socket packet buffer (ring buffer Linux kernel).
    - Circular logging system (menyimpan $N$ log terakhir dalam memori tanpa kebocoran RAM).

---

- [x] **ISSUE-09: Tree (Binary Tree & Binary Search Tree)**
  - **Deskripsi**: Struktur data hierarki non-linear dengan relasi parent-child.
  - **Materi**:
    - Konsep Tree: Root, Leaf, Height, Depth, Subtree.
    - Binary Search Tree (BST) property: `Left < Root < Right`.
    - Operasi: `Insert`, `Search`, `Delete`, `Min`, `Max`.
    - Tree Traversal: In-Order, Pre-Order, Post-Order, dan Level-Order (BFS).
  - **Deliverable**:
    - [x] `03_data_structures/07_tree/bst.go`
    - [x] `03_data_structures/07_tree/bst_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Document Object Model (DOM) pada rendering browser (HTML tree).
    - File system struktur folder/direktori.
    - Fondasi indexing database (B-Tree, B+ Tree di MySQL/Postgres).
    - Routing table di HTTP router (misalnya Radix Tree di Gin / Echo framework).

---

### Milestone 3: Algoritma Pengurutan (Sorting Algorithms)

- [x] **ISSUE-10: Bubble Sort**
  - **Deskripsi**: Algoritma pengurutan perbandingan bertetangga.
  - **Materi**:
    - Mekanisme perbandingan berpasangan dan penukaran (*swap*).
    - Kompleksitas: $O(N^2)$ worst/average, optimasi early exit $O(N)$ jika sudah terurut.
    - Karakteristik: Stable, In-Place.
  - **Deliverable**:
    - [x] `04_sorting/01_bubble_sort/bubble_sort.go`
    - [x] `04_sorting/01_bubble_sort/bubble_sort_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Media pembelajaran konsep dasar pertukaran data & algoritma.
    - Sistem mikrokontroler ultra-low resource jika ukuran data sangat kecil ($N < 10$) dan memori instruksi sangat terbatas.

---

- [x] **ISSUE-11: Selection Sort**
  - **Deskripsi**: Mengurutkan dengan mencari elemen minimum dan menempatkannya di posisi terdepan.
  - **Materi**:
    - Mekanisme pembagian array menjadi bagian terurut (*sorted part*) dan belum terurut (*unsorted part*).
    - Kompleksitas: Selalu $O(N^2)$ (tidak peduli data sudah terurut atau belum).
    - Karakteristik: Unstable (secara default), In-Place, jumlah swap maksimal $O(N)$.
  - **Deliverable**:
    - [x] `04_sorting/02_selection_sort/selection_sort.go`
    - [x] `04_sorting/02_selection_sort/selection_sort_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Berguna ketika biaya penulisan ke memori fisik (Flash Memory / EEPROM write cycle) jauh lebih mahal daripada membaca memori, karena jumlah swap dibatasi maksimal $N$ kali.

---

- [x] **ISSUE-12: Insertion Sort**
  - **Deskripsi**: Mengurutkan dengan cara mengambil satu per satu elemen dan menyisipkannya pada posisi yang tepat di sub-array terurut.
  - **Materi**:
    - Cara kerja seperti menyusun kartu di tangan.
    - Kompleksitas: $O(N^2)$ worst/average, $O(N)$ best case jika data hampir terurut.
    - Karakteristik: Stable, In-Place, adaptif.
  - **Deliverable**:
    - [x] `04_sorting/03_insertion_sort/insertion_sort.go`
    - [x] `04_sorting/03_insertion_sort/insertion_sort_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Mengurutkan data streaming real-time yang masuk sedikit demi sedikit.
    - Sub-rutin pada algoritma hybrid canggih (seperti Timsort di Python/Java dan Introsort/pdqsort di Go standard library) ketika ukuran partisi data sudah kecil ($N \le 16$ hingga $32$).

---

- [x] **ISSUE-13: Merge Sort**
  - **Deskripsi**: Algoritma pengurutan berbasis strategi *Divide and Conquer*.
  - **Materi**:
    - Pembagian rekursif menjadi sub-array hingga ukuran 1, lalu penggabungan terurut (*merge*).
    - Kompleksitas: Selalu $O(N \log N)$ (Best, Average, Worst).
    - Ruang tambahan (*Space Complexity*): $O(N)$.
    - Karakteristik: Stable, Out-of-Place (untuk array).
  - **Deliverable**:
    - [x] `04_sorting/04_merge_sort/merge_sort.go`
    - [x] `04_sorting/04_merge_sort/merge_sort_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - **External Sorting**: Mengurutkan file data raksasa (puluhan Gigabyte/Terabyte) yang tidak muat seluruhnya di RAM komputer (dibagi ke disk chunks lalu di-merge).
    - Pengurutan pada Linked List (efisien karena tidak memerlukan memori kontinu).

---

- [x] **ISSUE-14: Quick Sort**
  - **Deskripsi**: Algoritma *Divide and Conquer* berbasis pemilihan elemen poros (*pivot*).
  - **Materi**:
    - Strategi pemilihan pivot (First, Last, Random, Median of Three).
    - Skema partisi: Lomuto Partition vs Hoare Partition.
    - Kompleksitas: $O(N \log N)$ average, $O(N^2)$ worst case (saat pivot buruk).
    - Karakteristik: In-Place, Unstable.
  - **Deliverable**:
    - [x] `04_sorting/05_quick_sort/quick_sort.go`
    - [x] `04_sorting/05_quick_sort/quick_sort_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Fondasi algoritma sorting bawaan sistem operasi dan library utama (pdqsort pada `slices.Sort` di Go modern).
    - Pengurutan dalam pemrosesan grafika dan engine game karena cache-friendly dan in-place.

---

### Milestone 4: Algoritma Pencarian (Searching Algorithms)

- [x] **ISSUE-15: Linear Search**
  - **Deskripsi**: Pencarian berurutan dari elemen pertama hingga elemen terakhir.
  - **Materi**:
    - Mekanisme iterasi sekuensial.
    - Kondisi: Data tidak perlu terurut (*unsorted*).
    - Kompleksitas: $O(N)$ worst/average, $O(1)$ best case.
  - **Deliverable**:
    - [x] `05_searching/01_linear_search/linear_search.go`
    - [x] `05_searching/01_linear_search/linear_search_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Mencari elemen pada array kecil atau data yang tidak terstruktur/belum disortir.
    - Pencarian sederhana pada konfigurasi file kecil atau argumen CLI.

---

- [x] **ISSUE-16: Binary Search**
  - **Deskripsi**: Pencarian cepat dengan membagi dua area pencarian setiap iterasi.
  - **Materi**:
    - Prasyarat: Data **wajib terurut** (*sorted*).
    - Menghindari integer overflow saat menghitung titik tengah: `mid = low + (high - low) / 2`.
    - Kompleksitas: $O(\log N)$ time, $O(1)$ space.
  - **Deliverable**:
    - [x] `05_searching/02_binary_search/binary_search.go`
    - [x] `05_searching/02_binary_search/binary_search_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Fungsi pencarian standar: `sort.Search` / `slices.BinarySearch` di Go.
    - Pencarian index node pada database B-Tree.
    - Fitur `git bisect` untuk mencari commit awal yang menyebabkan regresi bug dalam git history.

---

- [x] **ISSUE-17: Interpolation Search**
  - **Deskripsi**: Perbaikan dari Binary Search untuk data terurut yang terdistribusi secara seragam (*uniformly distributed*).
  - **Materi**:
    - Rumus penentuan posisi target berdasarkan estimasi matematis:
      $$\text{pos} = \text{low} + \frac{(\text{target} - \text{arr}[\text{low}]) \times (\text{high} - \text{low})}{\text{arr}[\text{high}] - \text{arr}[\text{low}]}$$
    - Kompleksitas: $O(\log \log N)$ average pada data seragam, $O(N)$ worst case.
  - **Deliverable**:
    - [x] `05_searching/03_interpolation_search/interpolation_search.go`
    - [x] `05_searching/03_interpolation_search/interpolation_search_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Pencarian pada kamus digital atau buku telepon (manusia membuka halaman berdasarkan huruf abjad, bukan selalu membelah tepat di tengah).
    - Pencarian data sensor/telemetri dengan kenaikan nilai linear atau timestamp bertahap seragam.

---

- [x] **ISSUE-18: Jump Search**
  - **Deskripsi**: Pencarian pada data terurut dengan melompati sejumlah elemen tetap lalu melakukan linear search.
  - **Materi**:
    - Menentukan ukuran langkah optimal: $m = \sqrt{N}$.
    - Kompleksitas: $O(\sqrt{N})$ time, $O(1)$ space.
    - Posisi komparasi: Lebih cepat dari Linear Search, lebih lambat dari Binary Search, namun hanya memerlukan lompatan satu arah ke depan (*forward only*).
  - **Deliverable**:
    - [x] `05_searching/04_jump_search/jump_search.go`
    - [x] `05_searching/04_jump_search/jump_search_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Sistem media penyimpanan fisik di mana operasi melompat mundur (*backward jump/seek*) sangat lambat atau mahal dibanding melompat maju.
    - Pencarian pada linked list unrolled yang terindeks sebagian.

---

- [x] **ISSUE-19: Exponential Search**
  - **Deskripsi**: Pencarian rentang dengan melipatgandakan eksponen indeks ($1, 2, 4, 8, \dots$) kemudian mengeksekusi Binary Search di rentang yang ditemukan.
  - **Materi**:
    - Tahap 1: Menemukan batas rentang $[2^{k-1}, 2^k]$ di mana target berada.
    - Tahap 2: Menjalankan Binary Search pada rentang tersebut.
    - Kompleksitas: $O(\log i)$ di mana $i$ adalah posisi target.
  - **Deliverable**:
    - [x] `05_searching/05_exponential_search/exponential_search.go`
    - [x] `05_searching/05_exponential_search/exponential_search_test.go`
    - [x] Penjelasan modul `README.md`.
  - **Penerapan Dunia Nyata**:
    - Pencarian pada array tak terbatas (*unbounded/infinite arrays*), seperti streaming log atau live socket data stream di mana ukuran akhir data tidak diketahui sebelumnya.
    - Sangat efisien ketika elemen target berada dekat dengan posisi awal array.

---

### Milestone 5: Evaluasi Akhir & Sintesis (Penutup)

- [x] **ISSUE-20: Rangkuman, Benchmarking Komparatif & Panduan Pemilihan DSA**
  - **Deskripsi**: Evaluasi holistik seluruh materi yang telah dipelajari.
  - **Materi**:
    - Benchmark gabungan performa Sorting ($N=100, N=1000, N=10000$).
    - Benchmark gabungan performa Searching pada berbagai pola data (seragam, acak, terurut).
    - Matriks ringkasan kompleksitas waktu & memori (*Cheat Sheet*).
    - Tips praktis: Kapan harus menggunakan struktur data/algoritma tertentu di dunia kerja.
    - Topik lanjutan yang disarankan (Graph, Hash Table, Heap/Priority Queue, Dynamic Programming).
  - **Deliverable**:
    - [x] `06_penutup/README.md`
    - [x] `06_penutup/benchmark_comparison_test.go`
  - **Penerapan Dunia Nyata**:
    - Kemampuan membuat keputusan arsitektur (*architectural trade-off decision*) yang matang saat *system design* dan *code review*.

---

## 🛠️ Panduan Eksekusi Tugas

1. Kerjakan modul per modul secara berurutan sesuai nomor issue.
2. Setiap kali menyelesaikan modul, jalankan pengujian:
   ```bash
   go test -v ./...
   ```
3. Lakukan benchmark jika diperlukan:
   ```bash
   go test -bench=. -benchmem ./...
   ```
4. Beri tanda centang `[x]` pada issue terkait di file ini setelah implementasi dan pengujian lulus.


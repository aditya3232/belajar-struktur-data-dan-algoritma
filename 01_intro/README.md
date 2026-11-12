# 01 - Pengantar Struktur Data dan Algoritma (Golang)

Selamat datang di modul pembelajaran **Struktur Data dan Algoritma (DSA)** menggunakan bahasa pemrograman **Go (Golang)**!

Repositori ini dirancang bukan hanya untuk memahami teori dasar atau menghafal implementasi algoritma untuk wawancara teknis (*technical interview*), melainkan untuk **memahami cara kerja internal (*internals*), efisiensi alokasi memori, karakteristik CPU cache, dan penerapannya di dunia nyata pada sistem berskala besar (*high-scale systems*)**.

---

## 🚀 Mengapa Go untuk Belajar Struktur Data & Algoritma?

Go adalah bahasa pemrograman yang dirancang untuk membangun sistem backend modern, infrastruktur cloud (*cloud-native* seperti Docker dan Kubernetes), serta microservices berperforma tinggi.

Karakteristik unik Go dalam implementasi DSA:
1. **Explicit Memory Model (Pointers)**:
   Go memberikan kontrol langsung terhadap pointer tanpa pointer arithmetic yang berbahaya seperti di C/C++. Anda bisa menentukan apakah sebuah data dialokasikan sebagai value (pada stack/in-line) atau referensi pointer (pada heap).
2. **Generics (`any`, `comparable`, `constraints.Ordered`)**:
   Mulai dari Go 1.18+, kita dapat menulis struktur data dan algoritma yang *type-safe* dan dapat digunakan kembali (*reusable*) untuk berbagai tipe data tanpa mengorbankan performa (*runtime reflection overhead*).
3. **Rich Standard Tooling**:
   Go memiliki paket pengujian kelas satu (`testing`) bawaan tanpa perlu framework eksternal. Anda dapat menjalankan unit test, race detector, hingga benchmark alokasi memori secara instan:
   ```bash
   go test -v ./...
   go test -bench=. -benchmem ./...
   ```
4. **Garbage Collection & Cache Locality Awareness**:
   Memahami struktur data di Go mengajarkan Anda bagaimana mengurangi beban Garbage Collector (GC) dan memanfaatkan CPU cache locality (L1/L2/L3) secara optimal.

---

## 📂 Peta Modul Pembelajaran

```text
├── 01_intro/                 <- Filosofi & Setup
├── 02_overview_dsa/          <- Asymptotic Analysis (Big-O, Big-Omega, Big-Theta)
├── 03_data_structures/       <- Koleksi Data Dasar hingga Lanjutan
│   ├── 01_array_slice/
│   ├── 02_linked_list/
│   ├── 03_double_linked_list/
│   ├── 04_stack/
│   ├── 05_queue/
│   ├── 06_circular_queue/
│   └── 07_tree/
├── 04_sorting/               <- Algoritma Pengurutan
│   ├── 01_bubble_sort/
│   ├── 02_selection_sort/
│   ├── 03_insertion_sort/
│   ├── 04_merge_sort/
│   └── 05_quick_sort/
├── 05_searching/             <- Algoritma Pencarian
│   ├── 01_linear_search/
│   ├── 02_binary_search/
│   ├── 03_interpolation_search/
│   ├── 04_jump_search/
│   └── 05_exponential_search/
└── 06_penutup/               <- Sintesis, Cheat Sheet & Benchmark Komparatif
```

---

## 🛠️ Cara Menjalankan Kode dan Pengujian

### 1. Menjalankan Semua Unit Test
```bash
go test -v ./...
```

### 2. Menjalankan Test Pada Modul Tertentu
Contoh pada modul Stack:
```bash
go test -v ./03_data_structures/04_stack/...
```

### 3. Menjalankan Benchmark Alokasi Memori
```bash
go test -bench=. -benchmem ./...
```

---

## 🌐 Penerapan Dunia Nyata (Real-World Relevance)
Di industri teknologi, struktur data dan algoritma bukanlah soal menyelesaikan teka-teki logika semata. Ia adalah fondasi dari:
- Mengurangi biaya tagihan cloud (AWS/GCP/Azure) dengan meminimalkan penggunaan vCPU dan RAM.
- Mempercepat latensi respon API dari hitungan ratusan milidetik menjadi sub-milidetik.
- Mencegah insiden *out-of-memory* (OOM) dan sistem *crash* saat lonjakan traffic tinggi (*high concurrency / flash sales*).

# 03.4 - Stack (Tumpukan - LIFO)

Stack adalah struktur data linear yang bekerja dengan prinsip **LIFO (Last In, First Out)**: elemen yang terakhir dimasukkan adalah elemen yang pertama kali dikeluarkan.

```text
       ┌───────────┐
Push ─►│ Elemen 3  │──► Pop (Paling atas / Top)
       ├───────────┤
       │ Elemen 2  │
       ├───────────┤
       │ Elemen 1  │ (Paling bawah / Base)
       └───────────┘
```

---

## ⚙️ Operasi Utama & Kompleksitas

- **`Push(val)`**: Menambahkan elemen ke atas tumpukan ($O(1)$).
- **`Pop()`**: Mengambil dan menghapus elemen teratas ($O(1)$).
- **`Peek()`**: Melihat elemen teratas tanpa menghapusnya ($O(1)$).
- **`IsEmpty()`**: Memeriksa apakah tumpukan kosong ($O(1)$).
- **`Size()`**: Mengetahui jumlah elemen ($O(1)$).

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

### 1. Call Stack Eksekusi Fungsi di Runtime Go
Setiap kali fungsi memanggil fungsi lain (`A() -> B() -> C()`), Go runtime membuat *stack frame* baru di goroutine stack untuk menyimpan variabel lokal dan return address. Saat `C()` selesai, framenya di-*pop* dari stack.

### 2. Validasi Tanda Kurung & Sintaks Kode (Compiler/Linter)
Compiler (seperti `go build` atau Babel) menggunakan stack untuk memeriksa apakah pasangan kurung `()`, `{}`, dan `[]` seimbang dan bersarang dengan benar.

### 3. Operasi Undo / Redo di Software Produktivitas
Setiap aksi pengguna (misalnya mengetik teks atau mengubah warna shape) di-*push* ke tumpukan Undo. Menekan `Ctrl+Z` melakukan *pop* terhadap aksi terakhir.

### 4. Algoritma Backtracking & Depth-First Search (DFS)
Pencarian rute labirin, traversal pohon rekursif, atau penyelesaian teka-teki Sudoku menggunakan stack untuk mencatat jalur yang sedang ditelusuri.

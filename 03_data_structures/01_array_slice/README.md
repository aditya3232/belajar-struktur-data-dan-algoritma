# 03.1 - Array dan Slice di Go

Array dan Slice adalah struktur data linear paling mendasar dalam pemrograman Go. Memahami perbedaan keduanya dan bagaimana memori dikelola adalah kunci untuk menghasilkan aplikasi Go yang efisien dan minim alokasi memori.

---

## 📌 Array vs Slice: Perbedaan Fundamental

### 1. Array (Fixed Size, Value Type)
- Ukuran array ditentukan saat kompilasi dan menjadi bagian dari tipe data: `[4]int` berbeda tipe dengan `[5]int`.
- Array di Go adalah **value type**. Saat Anda mengoper array ke fungsi, Go menyalin seluruh isi array (bukan hanya pointer-nya).

```go
var arr [3]int = [3]int{1, 2, 3}
```

### 2. Slice (Dynamic, Reference Structure)
Slice adalah lapisan abstraksi (*view*) di atas array yang mendasarinya (*underlying array*). Di memori, sebuah slice direpresentasikan oleh struktur **Slice Header** (24 byte pada mesin 64-bit):
```text
┌───────────────────────────┐
│ Pointer (uintptr)         │ ───► Menunjuk ke elemen pertama underlying array
├───────────────────────────┤
│ Length (int)              │ ───► Jumlah elemen yang saat ini ada dalam slice
├───────────────────────────┤
│ Capacity (int)            │ ───► Total kapasitas underlying array dari titik pointer
└───────────────────────────┘
```

---

## ⚡ Cache Locality & Performa CPU

Array dan slice menyimpan elemen-elemennya pada **blok memori yang berdekatan (*contiguous memory block*)**.

Mengapa ini sangat cepat?
- **CPU Cache Line**: Ketika CPU membaca `slice[0]`, CPU tidak hanya mengambil 1 integer, melainkan memuat seluruh baris cache (biasanya 64 byte cache line) ke dalam CPU Cache L1/L2.
- **Cache Hit vs Cache Miss**: Mengakses elemen berikutnya (`slice[1]`, `slice[2]`) menghasilkan *cache hit*, yang 100x lebih cepat daripada mengambil data dari RAM utama. Struktur berbasis pointer (seperti Linked List) sering kali tersebar di heap dan memicu banyak *cache miss*.

---

## ⚠️ Waspadai Kebocoran Memori (Slice Memory Leak)

Jika Anda memiliki slice raksasa (misalnya 100 MB data file) dan hanya memotong 2 byte:
```go
hugeSlice := readHugeFile() // 100 MB
smallPiece := hugeSlice[:2] // Hanya butuh 2 byte!
```
Karena `smallPiece` masih mereferensikan underlying array 100 MB tersebut, **Garbage Collector tidak dapat membersihkan memori 100 MB itu dari RAM!**

Solusinya adalah menyalin data yang dibutuhkan ke slice baru:
```go
smallPiece := make([]byte, 2)
copy(smallPiece, hugeSlice[:2])
// Sekarang hugeSlice bisa dibersihkan oleh GC
```

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

1. **Buffer I/O Streaming (`bytes.Buffer` / `bufio.Reader`)**:
   Membaca paket data jaringan TCP atau file disk dalam potongan (*chunks*) byte tanpa overhead realokasi.
2. **Pemrosesan Gambar & Matriks Pixel**:
   Representasi gambar digital (RGBA) sebagai slice 1D kontinu yang dimanipulasi dengan perhitungan indeks `(y * width + x) * 4`.
3. **Database Batch Fetcher**:
   Mengambil ribuan record baris dari database ke dalam pre-allocated slice `make([]User, 0, batchSize)` untuk menghindari GC churn.

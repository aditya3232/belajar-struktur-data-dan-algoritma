# 04.2 - Selection Sort

Selection Sort adalah algoritma pengurutan sederhana yang bekerja dengan cara berulang kali mencari elemen terkecil (*minimum*) dari bagian array yang belum terurut, lalu menukarnya ke posisi terdepan sub-array tersebut.

```text
Awal:     [29, 10, 14, 37, 13]
Pass 1:   Cari min dari index 0..4 -> Ditemukan 10 (idx 1). Swap dengan arr[0]:
          [10, | 29, 14, 37, 13]
Pass 2:   Cari min dari index 1..4 -> Ditemukan 13 (idx 4). Swap dengan arr[1]:
          [10, 13, | 14, 37, 29]
Pass 3:   Cari min dari index 2..4 -> Ditemukan 14 (idx 2). Tetap:
          [10, 13, 14, | 37, 29]
Pass 4:   Cari min dari index 3..4 -> Ditemukan 29 (idx 4). Swap dengan arr[3]:
          [10, 13, 14, 29, | 37]
Selesai:  [10, 13, 14, 29, 37]
```

---

## ⚡ Karakteristik & Kompleksitas

- **Kompleksitas Waktu**:
  - **Best Case**: $O(N^2)$ (tetap harus memindai seluruh elemen untuk memastikan nilai minimum).
  - **Average Case**: $O(N^2)$
  - **Worst Case**: $O(N^2)$
- **Kompleksitas Ruang**: $O(1)$ (In-Place).
- **Jumlah Operasi Tulis (Memory Writes/Swaps)**: Maksimal $N - 1$ kali swap ($O(N)$ writes).

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

Meskipun $O(N^2)$ dalam perbandingan waktu, Selection Sort memiliki properti unik: **jumlah penulisan ke memori (memory write/swap) dibatasi secara ketat maksimal $N$ kali**.

- **EEPROM & Flash Memory Wear-Leveling**:
  Pada hardware mikrokontroler atau sensor IoT industri, memori non-volatile (Flash / EEPROM) memiliki batasan siklus penulisan fisik (misal maksimal 10.000 atau 100.000 kali write sebelum sektor rusak permanen) dan proses write membutuhkan energi baterai tinggi.
  Jika data perlu diurutkan langsung di flash memory, Selection Sort meminimalkan jumlah operasi write dibandingkan algoritma lain seperti Bubble Sort yang bisa melakukan hingga $O(N^2)$ write swaps.

# 05.1 - Linear Search (Pencarian Linear)

Linear Search (atau *Sequential Search*) adalah algoritma pencarian paling sederhana yang memeriksa setiap elemen dalam koleksi satu per satu dari awal hingga akhir sampai elemen yang dicari ditemukan atau seluruh koleksi selesai dipindai.

```text
Data:   [10, 50, 30, 70, 80, 20], Target: 70
Index 0: 10 == 70? (False)
Index 1: 50 == 70? (False)
Index 2: 30 == 70? (False)
Index 3: 70 == 70? (True -> Return Index 3)
```

---

## ⚡ Karakteristik & Analisis Kompleksitas

- **Prasyarat**: **Tidak ada**. Data tidak harus terurut (*can work on unsorted arrays*).
- **Kompleksitas Waktu**:
  - **Best Case**: $O(1)$ (target berada di indeks pertama).
  - **Average Case**: $O(N)$
  - **Worst Case**: $O(N)$ (target berada di akhir atau tidak ada).
- **Kompleksitas Ruang**: $O(1)$.

---

## 🏢 Kapan Linear Search Lebih Baik daripada Binary Search?

Meskipun Binary Search memiliki kompleksitas $O(\log N)$, ada situasi di mana Linear Search adalah pilihan terbaik:
1. **Data Berukuran Kecil ($N \le 30-50$)**:
   Jika Anda memiliki slice kecil, overhead algoritma pengurutan ($O(N \log N)$) agar bisa melakukan binary search jauh lebih mahal daripada langsung melakukan scan $O(N)$ satu kali.
2. **Data Dinamis yang Sering Ditulis Tapi Jarang Dicari**:
   Jika sistem sering melakukan *insert* data baru tanpa henti, memelihara array agar selalu terurut memakan biaya tinggi.
3. **Pencarian Flag / Argumen CLI Sederhana**:
   Memeriksa apakah flag seperti `--verbose` atau `--port` ada dalam slice `os.Args`.

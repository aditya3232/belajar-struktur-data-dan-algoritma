# 05.4 - Jump Search (Pencarian Lompat)

Jump Search adalah algoritma pencarian pada data terurut yang berada di antara Linear Search ($O(N)$) dan Binary Search ($O(\log N)$).

Alih-alih memeriksa setiap elemen ($1, 2, 3, 4, \dots$) atau melompat bolak-balik dua arah seperti Binary Search, Jump Search **hanya melompat maju dalam satu arah** dengan ukuran langkah tetap ($m$), lalu melakukan pencarian linear singkat saat blok yang memuat target telah terlampaui.

```text
Array: [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610] (N = 16)
Ukuran langkah optimal: m = sqrt(16) = 4

Target = 55
Lompat 1: Cek indeks 3 (nilai: 2)   < 55 -> Lanjut lompat
Lompat 2: Cek indeks 7 (nilai: 13)  < 55 -> Lanjut lompat
Lompat 3: Cek indeks 11 (nilai: 89) > 55 -> Target ada di antara indeks [8..11]!

Pencarian linear mundur/maju di blok [8..11]:
Cek arr[8] = 21
Cek arr[9] = 34
Cek arr[10] = 55 (Ditemukan di indeks 10!)
```

---

## 🧮 Penentuan Langkah Optimal $m = \sqrt{N}$

Jika ukuran langkah adalah $m$:
- Jumlah lompatan maksimal: $N / m$
- Jumlah perbandingan linear maksimal dalam blok: $m - 1$
- Total fungsi biaya: $f(m) = \frac{N}{m} + m$

Untuk meminimalkan fungsi ini, kita turunkan terhadap $m$ dan setel ke $0$:
$$-\frac{N}{m^2} + 1 = 0 \implies m^2 = N \implies m = \sqrt{N}$$

Maka total kompleksitas waktu adalah:
$$O\left(\frac{N}{\sqrt{N}} + \sqrt{N}\right) = O(\sqrt{N})$$

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

Mengapa menggunakan Jump Search jika Binary Search $O(\log N)$ lebih cepat?
1. **Media Penyimpanan dengan Forward-Only Seek**:
   Pada beberapa sistem penyimpanan fisik (seperti magnetic optical tape, disk serial stream, atau kompresi run-length), operasi memutar kembali / melompat mundur (*backward seeking*) membutuhkan penataan ulang head disk mekanik yang sangat mahal.
   Binary Search sering melompat bolak-balik (`mid -> left -> right -> left`), sedangkan Jump Search **hanya bergerak maju** dan hanya melakukan navigasi lokal dalam 1 blok kecil.
2. **Pencarian pada Unrolled Linked List**:
   Struktur data linked list di mana setiap node menampung array kecil.

# 03.6 - Circular Queue (Ring Buffer)

Circular Queue (dikenal juga sebagai **Ring Buffer**) adalah varian antrean dengan kapasitas memori tetap (*fixed memory buffer*) di mana indeks ujung akhir (*tail/rear*) terhubung kembali secara logis ke indeks ujung awal (*head/front*).

```text
               [0] ──► Head
             ▲     ▼
            /       \
          [3]       [1]
            \       /
             ▲     ▼
               [2] ──► Tail
```

---

## 🧮 Cara Kerja: Aritmatika Modulo (% Capacity)

Daripada memindahkan elemen dalam memori fisik (yang memakan biaya $O(N)$), Circular Queue memajukan penunjuk indeks menggunakan sisa bagi (modulo):
```go
nextTail = (tail + 1) % capacity
nextHead = (head + 1) % capacity
```
Dengan formula ini, ketika `tail` mencapai indeks terakhir (`capacity - 1`), operasi berikutnya otomatis kembali ke indeks `0` asalkan buffer belum penuh.

---

## ⚡ Mengapa Ring Buffer Sangat Disukai di Industri?

1. **Nol Alokasi Memori Tambahan (*Zero Allocation*)**:
   Setelah buffer awal dialokasikan (`make([]T, capacity)`), tidak pernah ada memori baru yang dibuat di heap. Ini menghilangkan beban Garbage Collector (GC latency spike).
2. **Cache Locality Maksimal**:
   Data disimpan dalam satu blok array datar kontinu di RAM.
3. **Penyangga Aliran Data Real-Time**:
   Jika produsen data menghasilkan data lebih cepat daripada konsumen memprosesnya, buffer bertindak sebagai penyerap lonjakan (*burst absorber*).

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

### 1. Audio & Video Streaming Engine
Ketika Anda menonton video di YouTube atau mendengarkan lagu di Spotify, data audio/video dari internet ditulis ke dalam Ring Buffer audio card. Pemutar suara membaca dari buffer ini secara konstan tanpa terputus (*anti-stuttering*).

### 2. Socket Network Buffer di Linux Kernel
Linux kernel menggunakan Ring Buffer pada driver kartu jaringan (NIC / Network Interface Card) seperti mekanisme `AF_PACKET ring buffer` untuk menampung paket TCP/UDP yang masuk dari kabel LAN sebelum diserahkan ke sistem operasi.

### 3. Circular Logging System (In-Memory Ring Log)
Menyimpan $N$ baris log debug terakhir di dalam memori microservice (seperti `dmesg` di Linux). Jika ada error, $N$ baris log terakhir dapat dicetak untuk inspeksi tanpa khawatir kehabisan RAM.

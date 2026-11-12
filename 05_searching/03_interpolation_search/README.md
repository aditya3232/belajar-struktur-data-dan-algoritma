# 05.3 - Interpolation Search (Pencarian Interpolasi)

Interpolation Search adalah evolusi dari Binary Search untuk data numerik terurut yang **terdistribusi secara seragam (*uniformly distributed*)**.

---

## 📖 Analogi Dunia Nyata: Membuka Buku Telepon / Kamus

Bayangkan Anda mencari nama seseorang berawalan **"Zul"** di dalam buku telepon tebal:
- **Binary Search**: Membuka tepat di halaman tengah (huruf "M"), lalu membelah lagi ke tengah (huruf "S"), membelah lagi ke tengah (huruf "V"), dst.
- **Interpolation Search**: Sebagai manusia cerdas, Anda tahu huruf "Z" berada sangat dekat dengan sampul belakang buku! Anda langsung membuka lembaran 95% ke arah akhir.

---

## 🧮 Rumus Estimasi Posisi

Formula interpolasi linear menghitung rasio jarak nilai target terhadap rentang nilai saat ini:
$$\text{pos} = \text{low} + \left[\frac{(\text{target} - \text{arr}[\text{low}]) \times (\text{high} - \text{low})}{\text{arr}[\text{high}] - \text{arr}[\text{low}]}\right]$$

---

## ⚡ Analisis Kompleksitas

- **Average Case**: **$O(\log \log N)$** (jika data terdistribusi rata / seragam).
  Sebagai contoh, untuk $N = 1.000.000$:
  - $\log_2(1.000.000) \approx 20$ langkah pada Binary Search.
  - $\log_2(\log_2(1.000.000)) \approx 4.3$ langkah pada Interpolation Search!
- **Worst Case**: **$O(N)$** (jika sebaran data sangat timpang atau eksponensial, misal: `[1, 2, 4, 8, 16, ..., 1000000]`).
- **Kompleksitas Ruang**: $O(1)$.

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

1. **Pencarian Data Telemetri & Sensor IoT**:
   Log sensor dengan sampling interval berkala (misal: suhu dicatat setiap 1 detik, menghasilkan timestamp `1001, 1002, 1003, ...`). Karena intervalnya hampir sempurna seragam, Interpolation Search menemukan timestamp target dalam 1-3 langkah saja.
2. **Lookup Table / Indeks Finansial Runtun Waktu (Time Series)**:
   Pencarian data tick harga saham pada interval milidetik reguler.

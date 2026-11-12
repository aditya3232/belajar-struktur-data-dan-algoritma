# 03.2 - Singly Linked List

Singly Linked List adalah struktur data linear di mana setiap elemen (*Node*) memiliki dua komponen utama:
1. **Value**: Data yang disimpan.
2. **Next Pointer**: Alamat memori yang menunjuk ke node berikutnya dalam urutan.

Node pertama disebut **Head**, dan node terakhir menunjuk ke `nil`.

```text
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│ Val: 10      │     │ Val: 20      │     │ Val: 30      │
│ Next: ───────┼────►│ Next: ───────┼────►│ Next: nil    │
└──────────────┘     └──────────────┘     └──────────────┘
     (Head)                                    (Tail)
```

---

## ⚖️ Linked List vs Array / Slice

| Karakteristik | Array / Slice | Singly Linked List |
|---|---|---|
| **Alokasi Memori** | Blok kontinu (*contiguous block*) | Tersebar di heap (*scattered nodes*) |
| **Cache Locality** | Sangat tinggi (CPU cache hit) | Rendah (banyak cache miss pointer chasing) |
| **Insert / Delete di Awal** | Lambat ($O(N)$ karena pergeseran) | Sangat cepat ($O(1)$ ubah pointer head) |
| **Akses Elemen acak (`idx`)**| Sangat cepat ($O(1)$) | Lambat ($O(N)$ harus traverse dari Head) |
| **Overhead Memori** | Sangat kecil (hanya data + slice header) | Ada overhead 8 byte per node untuk pointer `Next` |

---

## ⏱️ Kompleksitas Waktu Operasi

- **InsertHead**: $O(1)$
- **DeleteHead**: $O(1)$
- **InsertTail**: $O(N)$ (bisa $O(1)$ jika menyimpan pointer `tail`)
- **DeleteTail**: $O(N)$ (harus mencari node kedua terakhir)
- **Search**: $O(N)$
- **Reverse**: $O(N)$ waktu, $O(1)$ ruang in-place

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

1. **Resolusi Collision pada Hash Map (Separate Chaining)**:
   Ketika dua key berbeda menghasilkan hash index yang sama di dalam bucket Hash Table, Linked List digunakan untuk menampung kedua nilai tersebut pada bucket yang sama.
2. **Kernel Memory Allocator (Free-List)**:
   Sistem operasi melacak blok-blok memori RAM yang belum terpakai menggunakan linked list (*freelist*). Saat alokasi baru diminta, node diambil dari list.
3. **Undo History Sederhana**:
   Menyimpan snapshot aksi sebelumnya di mana setiap state menunjuk ke aksi sebelumnya.

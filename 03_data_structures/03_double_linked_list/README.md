# 03.3 - Doubly Linked List (Double Linked List)

Doubly Linked List adalah jenis linked list di mana setiap node memiliki **dua pointer**:
1. `Prev`: Menunjuk ke node sebelumnya.
2. `Next`: Menunjuk ke node setelahnya.

Selain itu, struktur ini umumnya memegang referensi langsung ke `head` dan `tail`.

```text
       ┌──────────────┐          ┌──────────────┐          ┌──────────────┐
nil ◄──┤ Prev: nil    │◄─────────┤ Prev: ───────┼─────────►│ Prev: ───────│
       │ Val:  10     │          │ Val:  20     │          │ Val:  30     │
       │ Next: ───────┼─────────►│ Next: ───────┼─────────►│ Next: nil    │──► nil
       └──────────────┘          └──────────────┘          └──────────────┘
            (Head)                                              (Tail)
```

---

## ⚡ Keunggulan Doubly Linked List

1. **Penghapusan Node $O(1)$ Instan**:
   Pada Singly Linked List, menghapus node tertentu membutuhkan pencarian node *sebelumnya* ($O(N)$). Pada Doubly Linked List, karena node sudah memiliki pointer `node.Prev`, pelepasan sambungan (*unlinking*) langsung dapat dilakukan dalam $O(1)$:
   ```go
   node.Prev.Next = node.Next
   node.Next.Prev = node.Prev
   ```
2. **Penambahan dan Pengurangan di Kedua Ujung $O(1)$**:
   Operasi `PushFront`, `PushBack`, `PopFront`, dan `PopBack` semuanya berjalan dalam waktu konstan $O(1)$.
3. **Traversial Dua Arah**:
   Dapat dijelajahi dari depan ke belakang (*Forward*) maupun sebaliknya (*Backward*).

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

### 1. LRU Cache (Least Recently Used Cache)
Di backend database atau caching server (seperti Redis atau in-memory cache), LRU Cache adalah mekanisme penggusuran data terpopuler:
- Data diakses via **Hash Map** untuk lookup $O(1)$.
- Urutan penggunaan dilacak via **Doubly Linked List**:
  - Item yang baru dibaca/ditulis dipindahkan ke `head` (Most Recently Used).
  - Ketika kapasitas RAM penuh, item di posisi `tail` (Least Recently Used) dihapus dalam $O(1)$.

### 2. Tab & Riwayat Web Browser (Back & Forward)
Setiap halaman web yang Anda kunjungi disimpan dalam node. Tombol *Back* bergerak ke `node.Prev`, dan tombol *Forward* bergerak ke `node.Next`.

### 3. Playlist Media Player
Fitur memutar lagu berikutnya (*Next*) atau lagu sebelumnya (*Previous*), serta menyisipkan lagu ke dalam antrean lagu yang sedang berjalan.

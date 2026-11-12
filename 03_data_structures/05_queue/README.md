# 03.5 - Queue (Antrean - FIFO)

Queue adalah struktur data linear yang bekerja dengan prinsip **FIFO (First In, First Out)**: elemen yang pertama kali masuk akan menjadi elemen yang pertama kali keluar.

```text
               ┌───────────┬───────────┬───────────┐
Enqueue (Belakang) ──►│ Elemen 3  │ Elemen 2  │ Elemen 1  │──► Dequeue (Depan)
               └───────────┴───────────┴───────────┘
```

---

## ⚠️ Mengapa Tidak Menggunakan Slice Naif?

Implementasi queue menggunakan slice sederhana:
```go
// Enqueue
q = append(q, val) // O(1)

// Dequeue
val := q[0]
q = q[1:] // O(1) indeks, tapi memori elemen lama tetap tertahan di underlying array!
```
Atau jika menghapus dengan copy/shift:
```go
q = append(q[:0], q[1:]...) // O(N) karena harus menggeser seluruh elemen ke kiri!
```
Oleh karena itu, implementasi queue standar industri menggunakan:
1. **Linked-List dengan Head & Tail pointer** (digunakan pada modul ini untuk ukuran tak terbatas).
2. **Circular Queue / Ring Buffer** (untuk kapasitas memori tetap tanpa alokasi baru).

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

### 1. Message Broker & Asynchronous Task Queue
RabbitMQ, Apache Kafka consumer, AWS SQS, atau Redis Celery/Asynq memproses pekerjaan (seperti pengiriman email, encoding video, atau webhook call) berdasarkan urutan antrean FIFO.

### 2. Web Server Request Buffer & Connection Pool
Saat server web (misal NGINX atau Go `net/http`) menerima ribuan request bersamaan yang melebihi jumlah *worker goroutine*, request ditampung dalam queue agar tidak terjadi *request drop*.

### 3. Breadth-First Search (BFS)
Algoritma penjelajahan graf dan pohon tingkat demi tingkat (*level-by-level*) memerlukan Queue untuk menyimpan tetangga (*neighbors*) yang akan dikunjungi berikutnya.

# 03.7 - Tree & Binary Search Tree (BST)

Tree adalah struktur data hierarkis non-linear yang terdiri dari sekumpulan simpul (*nodes*) yang saling terhubung melalui relasi induk-anak (*parent-child*).

```text
               [50] ◄── Root (Tingkat 1)
             /      \
          [30]      [70] ◄── Internal Node (Tingkat 2)
         /    \    /    \
       [20]  [40][60]  [80] ◄── Leaves (Daun / Tingkat 3)
```

---

## 🌲 Aturan Binary Search Tree (BST Property)

Pada Binary Search Tree, untuk setiap node:
1. Semua nilai di **subtree kiri** harus **lebih kecil** dari nilai node tersebut (`Left < Node`).
2. Semua nilai di **subtree kanan** harus **lebih besar** dari nilai node tersebut (`Right > Node`).

Properti ini memungkinkan pencarian, penambahan, dan penghapusan berjalan dalam **$O(\log N)$** pada pohon yang seimbang (*balanced*).

---

## 🔄 Strategi Penjelajahan Pohon (Tree Traversals)

1. **Depth-First Search (DFS)**:
   - **In-Order (Left -> Root -> Right)**: Menghasilkan deret angka yang selalu terurut secara menaik (*strictly increasing*).
   - **Pre-Order (Root -> Left -> Right)**: Berguna untuk menyalin atau membuat salinan kloning pohon (*serialization*).
   - **Post-Order (Left -> Right -> Root)**: Berguna saat menghapus pohon dari bawah ke atas (*bottom-up deletion*) atau menghitung ukuran direktori file.
2. **Breadth-First Search (BFS) / Level-Order**:
   - Menjelajahi pohon tingkat demi tingkat dari atas ke bawah menggunakan bantuan Queue.

---

## 🏢 Penerapan di Dunia Nyata (Real-World Applications)

### 1. Document Object Model (DOM Tree) di Web Browser
Halaman web HTML direpresentasikan sebagai pohon hierarkis (`<html> -> <body> -> <div> -> <p>`). Manipulasi elemen JavaScript (`document.getElementById`) menjelajahi pohon ini.

### 2. Struktur Direktori File System
Sistem berkas di Linux / Windows (misal `/home/user/documents/file.txt`) adalah pohon n-ary di mana folder adalah node induk dan berkas adalah daun (*leaf*).

### 3. Fondasi Indexing Database (B-Tree & B+ Tree)
Mesin database seperti MySQL (InnoDB) dan PostgreSQL menggunakan varian BST multi-cabang (B+ Tree) untuk mengindeks jutaan baris data agar query `WHERE id = 123` atau `WHERE id BETWEEN 10 AND 50` dapat dieksekusi dalam sub-milidetik.

### 4. Router HTTP URL Matching (Radix Tree)
Framework web Go modern seperti **Gin** atau **Echo** menggunakan Radix Tree (varian pohon pencarian prefix) untuk mencocokkan route URL (`/users/:id/posts`) dalam hitungan nanodetik.

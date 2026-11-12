package tree

import (
	"errors"
)

var (
	ErrEmptyTree = errors.New("tree kosong")
)

// Ordered adalah constraint untuk tipe data yang mendukung operator pembanding (<, <=, >, >=).
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Node merepresentasikan simpul dalam Binary Search Tree (BST).
type Node[T Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// BinarySearchTree menyimpan pointer ke root node dan melacak jumlah simpul.
type BinarySearchTree[T Ordered] struct {
	root *Node[T]
	size int
}

// NewBinarySearchTree menginisialisasi BST baru.
func NewBinarySearchTree[T Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{root: nil, size: 0}
}

// Size mengembalikan jumlah node dalam BST. O(1).
func (t *BinarySearchTree[T]) Size() int {
	return t.size
}

// IsEmpty memeriksa apakah tree kosong. O(1).
func (t *BinarySearchTree[T]) IsEmpty() bool {
	return t.size == 0
}

// Insert menyisipkan nilai baru ke dalam BST sesuai aturan Left < Root < Right.
// Kompleksitas Waktu: O(log N) average, O(N) worst case.
func (t *BinarySearchTree[T]) Insert(val T) {
	t.root = t.insertNode(t.root, val)
}

func (t *BinarySearchTree[T]) insertNode(curr *Node[T], val T) *Node[T] {
	if curr == nil {
		t.size++
		return &Node[T]{Value: val, Left: nil, Right: nil}
	}

	if val < curr.Value {
		curr.Left = t.insertNode(curr.Left, val)
	} else if val > curr.Value {
		curr.Right = t.insertNode(curr.Right, val)
	}
	// Jika val == curr.Value (duplikat), kita abaikan untuk menjaga keunikan set
	return curr
}

// Search mencari nilai tertentu di dalam BST.
// Kompleksitas Waktu: O(log N) average, O(N) worst case.
func (t *BinarySearchTree[T]) Search(val T) bool {
	curr := t.root
	for curr != nil {
		if val == curr.Value {
			return true
		} else if val < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return false
}

// Min mengembalikan nilai terkecil di dalam BST (elemen paling kiri). O(log N).
func (t *BinarySearchTree[T]) Min() (T, error) {
	var zero T
	if t.IsEmpty() {
		return zero, ErrEmptyTree
	}

	curr := t.root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Value, nil
}

// Max mengembalikan nilai terbesar di dalam BST (elemen paling kanan). O(log N).
func (t *BinarySearchTree[T]) Max() (T, error) {
	var zero T
	if t.IsEmpty() {
		return zero, ErrEmptyTree
	}

	curr := t.root
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr.Value, nil
}

// Delete menghapus nilai dari BST dan menata ulang relasi parent-child.
func (t *BinarySearchTree[T]) Delete(val T) {
	var deleted bool
	t.root, deleted = t.deleteNode(t.root, val)
	if deleted {
		t.size--
	}
}

func (t *BinarySearchTree[T]) deleteNode(curr *Node[T], val T) (*Node[T], bool) {
	if curr == nil {
		return nil, false
	}

	var deleted bool
	if val < curr.Value {
		curr.Left, deleted = t.deleteNode(curr.Left, val)
	} else if val > curr.Value {
		curr.Right, deleted = t.deleteNode(curr.Right, val)
	} else {
		// Ditemukan node yang akan dihapus:
		deleted = true

		// Kasus 1: Leaf node (tanpa anak)
		if curr.Left == nil && curr.Right == nil {
			return nil, true
		}

		// Kasus 2: Hanya memiliki 1 anak
		if curr.Left == nil {
			return curr.Right, true
		}
		if curr.Right == nil {
			return curr.Left, true
		}

		// Kasus 3: Memiliki 2 anak
		// Cari in-order successor (nilai terkecil di subtree kanan)
		successor := curr.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		curr.Value = successor.Value
		// Hapus in-order successor dari subtree kanan
		curr.Right, _ = t.deleteNode(curr.Right, successor.Value)
	}

	return curr, deleted
}

// InOrder mengembalikan urutan transversal Left -> Root -> Right (selalu terurut menaik). O(N).
func (t *BinarySearchTree[T]) InOrder() []T {
	result := make([]T, 0, t.size)
	t.inOrderHelper(t.root, &result)
	return result
}

func (t *BinarySearchTree[T]) inOrderHelper(curr *Node[T], res *[]T) {
	if curr == nil {
		return
	}
	t.inOrderHelper(curr.Left, res)
	*res = append(*res, curr.Value)
	t.inOrderHelper(curr.Right, res)
}

// PreOrder mengembalikan urutan Root -> Left -> Right (cocok untuk kloning/serialisasi pohon). O(N).
func (t *BinarySearchTree[T]) PreOrder() []T {
	result := make([]T, 0, t.size)
	t.preOrderHelper(t.root, &result)
	return result
}

func (t *BinarySearchTree[T]) preOrderHelper(curr *Node[T], res *[]T) {
	if curr == nil {
		return
	}
	*res = append(*res, curr.Value)
	t.preOrderHelper(curr.Left, res)
	t.preOrderHelper(curr.Right, res)
}

// PostOrder mengembalikan urutan Left -> Right -> Root (cocok untuk pembersihan memori bottom-up). O(N).
func (t *BinarySearchTree[T]) PostOrder() []T {
	result := make([]T, 0, t.size)
	t.postOrderHelper(t.root, &result)
	return result
}

func (t *BinarySearchTree[T]) postOrderHelper(curr *Node[T], res *[]T) {
	if curr == nil {
		return
	}
	t.postOrderHelper(curr.Left, res)
	t.postOrderHelper(curr.Right, res)
	*res = append(*res, curr.Value)
}

// LevelOrder mengembalikan transversal tingkat demi tingkat (BFS traversal). O(N).
func (t *BinarySearchTree[T]) LevelOrder() []T {
	if t.root == nil {
		return []T{}
	}

	result := make([]T, 0, t.size)
	queue := []*Node[T]{t.root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		result = append(result, curr.Value)
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return result
}

// Height menghitung tinggi maksimum pohon (level terdalam).
func (t *BinarySearchTree[T]) Height() int {
	return t.heightHelper(t.root)
}

func (t *BinarySearchTree[T]) heightHelper(curr *Node[T]) int {
	if curr == nil {
		return 0
	}
	leftHeight := t.heightHelper(curr.Left)
	rightHeight := t.heightHelper(curr.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

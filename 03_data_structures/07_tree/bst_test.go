package tree

import (
	"reflect"
	"testing"
)

func TestBST_BasicOps(t *testing.T) {
	bst := NewBinarySearchTree[int]()

	if !bst.IsEmpty() || bst.Size() != 0 {
		t.Errorf("BST baru harus kosong")
	}

	// Min & Max on empty tree
	_, err := bst.Min()
	if err != ErrEmptyTree {
		t.Errorf("Expected ErrEmptyTree on Min, got %v", err)
	}

	_, err = bst.Max()
	if err != ErrEmptyTree {
		t.Errorf("Expected ErrEmptyTree on Max, got %v", err)
	}

	/*
	         50
	       /    \
	     30      70
	    /  \    /  \
	   20  40  60  80
	*/
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		bst.Insert(v)
	}

	if bst.Size() != 7 {
		t.Errorf("Expected size 7, got %d", bst.Size())
	}

	// Search
	if !bst.Search(40) {
		t.Errorf("Search(40) should return true")
	}
	if bst.Search(999) {
		t.Errorf("Search(999) should return false")
	}

	// Min & Max
	minVal, _ := bst.Min()
	maxVal, _ := bst.Max()
	if minVal != 20 || maxVal != 80 {
		t.Errorf("Expected min 20 and max 80, got min %d, max %d", minVal, maxVal)
	}

	// InOrder Traversal (wajib selalu terurut menaik)
	inOrder := bst.InOrder()
	expectedInOrder := []int{20, 30, 40, 50, 60, 70, 80}
	if !reflect.DeepEqual(inOrder, expectedInOrder) {
		t.Errorf("InOrder expected %v, got %v", expectedInOrder, inOrder)
	}

	// PreOrder Traversal (Root -> Left -> Right)
	preOrder := bst.PreOrder()
	expectedPreOrder := []int{50, 30, 20, 40, 70, 60, 80}
	if !reflect.DeepEqual(preOrder, expectedPreOrder) {
		t.Errorf("PreOrder expected %v, got %v", expectedPreOrder, preOrder)
	}

	// PostOrder Traversal (Left -> Right -> Root)
	postOrder := bst.PostOrder()
	expectedPostOrder := []int{20, 40, 30, 60, 80, 70, 50}
	if !reflect.DeepEqual(postOrder, expectedPostOrder) {
		t.Errorf("PostOrder expected %v, got %v", expectedPostOrder, postOrder)
	}

	// LevelOrder Traversal (BFS)
	levelOrder := bst.LevelOrder()
	expectedLevelOrder := []int{50, 30, 70, 20, 40, 60, 80}
	if !reflect.DeepEqual(levelOrder, expectedLevelOrder) {
		t.Errorf("LevelOrder expected %v, got %v", expectedLevelOrder, levelOrder)
	}

	if bst.Height() != 3 {
		t.Errorf("Expected height 3, got %d", bst.Height())
	}
}

func TestBST_DeleteCases(t *testing.T) {
	/*
	         50
	       /    \
	     30      70
	    /  \       \
	   20  40       80
	*/
	bst := NewBinarySearchTree[int]()
	for _, v := range []int{50, 30, 70, 20, 40, 80} {
		bst.Insert(v)
	}

	// 1. Delete leaf node (20)
	bst.Delete(20)
	if bst.Search(20) || bst.Size() != 5 {
		t.Errorf("Failed to delete leaf node 20")
	}

	// 2. Delete node with 1 child (70, hanya punya anak kanan 80)
	bst.Delete(70)
	if bst.Search(70) || bst.Size() != 4 || !bst.Search(80) {
		t.Errorf("Failed to delete node with 1 child 70")
	}

	// 3. Delete node with 2 children (root 50, anak: 30 dan 80)
	bst.Delete(50)
	if bst.Search(50) || bst.Size() != 3 {
		t.Errorf("Failed to delete node with 2 children 50")
	}

	// Verifikasi BST tetap terurut
	expected := []int{30, 40, 80}
	if !reflect.DeepEqual(bst.InOrder(), expected) {
		t.Errorf("InOrder after deletions expected %v, got %v", expected, bst.InOrder())
	}
}

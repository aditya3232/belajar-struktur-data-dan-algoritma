package doublelinkedlist

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList_BasicOps(t *testing.T) {
	dll := NewDoublyLinkedList[int]()

	if !dll.IsEmpty() || dll.Size() != 0 {
		t.Errorf("List baru harus kosong")
	}

	dll.PushFront(20)
	dll.PushFront(10) // 10 <-> 20
	dll.PushBack(30)  // 10 <-> 20 <-> 30

	if dll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", dll.Size())
	}

	forward := dll.TraverseForward()
	expectedForward := []int{10, 20, 30}
	if !reflect.DeepEqual(forward, expectedForward) {
		t.Errorf("TraverseForward expected %v, got %v", expectedForward, forward)
	}

	backward := dll.TraverseBackward()
	expectedBackward := []int{30, 20, 10}
	if !reflect.DeepEqual(backward, expectedBackward) {
		t.Errorf("TraverseBackward expected %v, got %v", expectedBackward, backward)
	}

	// PopFront
	val, err := dll.PopFront()
	if err != nil || val != 10 {
		t.Errorf("PopFront expected 10, got %d, err %v", val, err)
	}

	// PopBack
	val, err = dll.PopBack()
	if err != nil || val != 30 {
		t.Errorf("PopBack expected 30, got %d, err %v", val, err)
	}

	if dll.Size() != 1 {
		t.Errorf("Expected size 1 after pops, got %d", dll.Size())
	}
}

func TestDoublyLinkedList_RemoveNode(t *testing.T) {
	dll := NewDoublyLinkedList[string]()
	n1 := dll.PushBack("A")
	n2 := dll.PushBack("B")
	n3 := dll.PushBack("C")

	// Hapus node tengah ("B")
	dll.RemoveNode(n2)
	expected := []string{"A", "C"}
	if !reflect.DeepEqual(dll.TraverseForward(), expected) {
		t.Errorf("Remove middle node expected %v, got %v", expected, dll.TraverseForward())
	}

	// Hapus head ("A")
	dll.RemoveNode(n1)
	expected = []string{"C"}
	if !reflect.DeepEqual(dll.TraverseForward(), expected) {
		t.Errorf("Remove head node expected %v, got %v", expected, dll.TraverseForward())
	}

	// Hapus tail ("C")
	dll.RemoveNode(n3)
	if !dll.IsEmpty() {
		t.Errorf("List should be empty after removing all nodes")
	}
}

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache[string, int](2)

	cache.Put("item1", 100)
	cache.Put("item2", 200)

	val, err := cache.Get("item1")
	if err != nil || val != 100 {
		t.Fatalf("Get('item1') expected 100, got %d, err %v", val, err)
	}

	// Masukkan item3, kapasitas penuh (2), item2 harus di-evict karena item1 baru saja di-akses (MRU)
	cache.Put("item3", 300)

	_, err = cache.Get("item2")
	if err != ErrKeyNotFound {
		t.Errorf("item2 should have been evicted by LRU policy")
	}

	val, err = cache.Get("item3")
	if err != nil || val != 300 {
		t.Errorf("item3 should be accessible")
	}
}

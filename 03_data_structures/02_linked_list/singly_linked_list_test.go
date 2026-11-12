package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList_BasicOps(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	if !list.IsEmpty() || list.Size() != 0 {
		t.Errorf("List baru harus kosong")
	}

	// InsertHead
	list.InsertHead(10)
	list.InsertHead(20) // List: 20 -> 10
	if list.Size() != 2 {
		t.Errorf("Size expected 2, got %d", list.Size())
	}

	// InsertTail
	list.InsertTail(30) // List: 20 -> 10 -> 30
	expected := []int{20, 10, 30}
	if !reflect.DeepEqual(list.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, list.ToSlice())
	}

	// InsertAt
	err := list.InsertAt(1, 15) // List: 20 -> 15 -> 10 -> 30
	if err != nil {
		t.Fatalf("InsertAt error: %v", err)
	}
	expected = []int{20, 15, 10, 30}
	if !reflect.DeepEqual(list.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, list.ToSlice())
	}

	// Search
	idx, found := list.Search(15)
	if !found || idx != 1 {
		t.Errorf("Search(15) expected (1, true), got (%d, %v)", idx, found)
	}

	_, found = list.Search(999)
	if found {
		t.Errorf("Search(999) should return false")
	}

	// DeleteHead
	val, err := list.DeleteHead() // Hapus 20. List: 15 -> 10 -> 30
	if err != nil || val != 20 {
		t.Errorf("DeleteHead expected 20, got %d, err: %v", val, err)
	}

	// DeleteTail
	val, err = list.DeleteTail() // Hapus 30. List: 15 -> 10
	if err != nil || val != 30 {
		t.Errorf("DeleteTail expected 30, got %d, err: %v", val, err)
	}

	// DeleteValue
	deleted := list.DeleteValue(15) // Hapus 15. List: 10
	if !deleted || list.Size() != 1 {
		t.Errorf("DeleteValue(15) failed")
	}

	deleted = list.DeleteValue(888)
	if deleted {
		t.Errorf("DeleteValue(888) should be false")
	}
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	list.Reverse() // Reverse list kosong tidak boleh panic
	if list.Size() != 0 {
		t.Errorf("Empty reverse size != 0")
	}

	list.InsertTail("A")
	list.Reverse() // Reverse 1 elemen
	if !reflect.DeepEqual(list.ToSlice(), []string{"A"}) {
		t.Errorf("Single element reverse failed")
	}

	list.InsertTail("B")
	list.InsertTail("C") // A -> B -> C
	list.Reverse()        // C -> B -> A

	expected := []string{"C", "B", "A"}
	if !reflect.DeepEqual(list.ToSlice(), expected) {
		t.Errorf("Reverse expected %v, got %v", expected, list.ToSlice())
	}
}

func TestSinglyLinkedList_EdgeCases(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	_, err := list.DeleteHead()
	if err != ErrEmptyList {
		t.Errorf("Expected ErrEmptyList on DeleteHead")
	}

	_, err = list.DeleteTail()
	if err != ErrEmptyList {
		t.Errorf("Expected ErrEmptyList on DeleteTail")
	}

	err = list.InsertAt(-1, 100)
	if err != ErrIndexOutOfRange {
		t.Errorf("Expected ErrIndexOutOfRange on index -1")
	}

	err = list.InsertAt(1, 100)
	if err != ErrIndexOutOfRange {
		t.Errorf("Expected ErrIndexOutOfRange on index > size")
	}
}

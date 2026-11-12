package queue

import (
	"reflect"
	"testing"
)

func TestQueue_BasicOps(t *testing.T) {
	q := NewQueue[int]()

	if !q.IsEmpty() || q.Size() != 0 {
		t.Errorf("Queue baru harus kosong")
	}

	// Dequeue & Front pada queue kosong
	_, err := q.Dequeue()
	if err != ErrEmptyQueue {
		t.Errorf("Expected ErrEmptyQueue on Dequeue, got %v", err)
	}

	_, err = q.Front()
	if err != ErrEmptyQueue {
		t.Errorf("Expected ErrEmptyQueue on Front, got %v", err)
	}

	// Enqueue elements
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	front, err := q.Front()
	if err != nil || front != 10 {
		t.Errorf("Front expected 10, got %d, err %v", front, err)
	}

	// Dequeue elements (FIFO order)
	val, err := q.Dequeue()
	if err != nil || val != 10 {
		t.Errorf("Dequeue expected 10, got %d", val)
	}

	val, err = q.Dequeue()
	if err != nil || val != 20 {
		t.Errorf("Dequeue expected 20, got %d", val)
	}

	val, err = q.Dequeue()
	if err != nil || val != 30 {
		t.Errorf("Dequeue expected 30, got %d", val)
	}

	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after dequeuing all elements")
	}
}

func TestSimpleTaskDispatcher(t *testing.T) {
	tasks := []Task{
		{ID: 101, Payload: "Kirim Email"},
		{ID: 102, Payload: "Resize Gambar"},
		{ID: 103, Payload: "Generate PDF"},
	}

	processed := SimpleTaskDispatcher(tasks)
	expected := []int{101, 102, 103}

	if !reflect.DeepEqual(processed, expected) {
		t.Errorf("Dispatcher FIFO order expected %v, got %v", expected, processed)
	}
}

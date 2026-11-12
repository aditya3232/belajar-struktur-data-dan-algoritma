package circularqueue

import (
	"reflect"
	"testing"
)

func TestCircularQueue_BasicOps(t *testing.T) {
	cq, err := NewCircularQueue[int](3)
	if err != nil {
		t.Fatalf("Failed to create circular queue: %v", err)
	}

	if !cq.IsEmpty() || cq.Size() != 0 || cq.IsFull() {
		t.Errorf("New circular queue must be empty and not full")
	}

	// Dequeue on empty
	_, err = cq.Dequeue()
	if err != ErrQueueEmpty {
		t.Errorf("Expected ErrQueueEmpty, got %v", err)
	}

	// Enqueue up to capacity
	_ = cq.Enqueue(10)
	_ = cq.Enqueue(20)
	_ = cq.Enqueue(30)

	if !cq.IsFull() || cq.Size() != 3 {
		t.Errorf("Queue should be full with size 3")
	}

	// Enqueue on full
	err = cq.Enqueue(40)
	if err != ErrQueueFull {
		t.Errorf("Expected ErrQueueFull, got %v", err)
	}

	// Front & Rear
	front, _ := cq.Front()
	rear, _ := cq.Rear()
	if front != 10 || rear != 30 {
		t.Errorf("Expected front 10 and rear 30, got front %d, rear %d", front, rear)
	}

	// Dequeue one element
	val, err := cq.Dequeue()
	if err != nil || val != 10 {
		t.Errorf("Dequeue expected 10, got %d", val)
	}

	// Wrap-around enqueue
	err = cq.Enqueue(40)
	if err != nil {
		t.Fatalf("Wrap-around enqueue failed: %v", err)
	}

	expected := []int{20, 30, 40}
	if !reflect.DeepEqual(cq.ToSlice(), expected) {
		t.Errorf("Wrap around expected %v, got %v", expected, cq.ToSlice())
	}
}

func TestCircularQueue_Overwrite(t *testing.T) {
	cq, _ := NewCircularQueue[string](3)

	cq.EnqueueOverwrite("log1")
	cq.EnqueueOverwrite("log2")
	cq.EnqueueOverwrite("log3")
	// Kapasitas 3 penuh: ["log1", "log2", "log3"]

	// Overwrite: "log1" akan dibuang digantikan "log4"
	cq.EnqueueOverwrite("log4")

	expected := []string{"log2", "log3", "log4"}
	if !reflect.DeepEqual(cq.ToSlice(), expected) {
		t.Errorf("Overwrite expected %v, got %v", expected, cq.ToSlice())
	}

	cq.EnqueueOverwrite("log5")
	expected = []string{"log3", "log4", "log5"}
	if !reflect.DeepEqual(cq.ToSlice(), expected) {
		t.Errorf("Overwrite expected %v, got %v", expected, cq.ToSlice())
	}
}

func TestCircularQueue_InvalidCapacity(t *testing.T) {
	_, err := NewCircularQueue[int](0)
	if err != ErrInvalidCapacity {
		t.Errorf("Expected ErrInvalidCapacity, got %v", err)
	}
}

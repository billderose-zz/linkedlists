package linked

import (
	"math/rand"
	"testing"
	"time"
)

func TestEmpty(t *testing.T) {

	list := New()

	if list.IsEmpty() {
		t.Log("Empty Passed")
	} else {
		t.Error("Error on Empty")
	}

	for i := 0; i < 1000; i++ {
		list.PushFront(i)
	}

	if !list.IsEmpty() {
		t.Log("Empty Passed")
	} else {
		t.Error("Error on Empty")
	}

	if list.Size() == 1000 {
		t.Log("PushFront Passed")
	} else {
		t.Error("PushFront Failed")
	}

	if e := list.Front().Value(); e == 999 {
		t.Log("Front Passed")
	} else {
		t.Error("Front Failed")
	}

	if e := list.Back().Value(); e == 0 {
		t.Log("Back Passed")
	} else {
		t.Error("Back Failed")
	}

	if e, _ := list.PeekFront(); e == 999 {
		t.Log("Front Passed")
	} else {
		t.Error("PeekFront Failed")
	}

	if e, _ := list.PeekBack(); e == 0 {
		t.Log("PeekBack Passed")
	} else {
		t.Error("PeekBack Failed")
	}

	if val, exists := list.Contains(95); exists && val.Value() == 95 {
		t.Log("Contains Passed")
	} else {
		t.Error("Contains Failed")
	}

	if val := list.Remove(89); val == 89 {
		t.Log("Remove Passed")
	} else {
		t.Error("Remove Failed")
	}

	list.Clear()

	for i := 0; i < 1000; i++ {
		list.PushBack(i)
	}

	if !list.IsEmpty() {
		t.Log("Empty Passed")
	} else {
		t.Error("Empty Failed")
	}

	if list.Size() == 1000 {
		t.Log("PushFront Passed")
	} else {
		t.Error("PushFront Failed")
	}

	if e := list.Front().Value(); e == 0 {
		t.Log("Front Passed")
	} else {
		t.Error("Front Failed")
	}

	if e := list.Back().Value(); e == 999 {
		t.Log("Back Passed")
	} else {
		t.Error("Back Failed")
	}

	if e, _ := list.PopFront(); e == 0 {
		t.Log("Front Passed")
	} else {
		t.Error("PeekFront Failed")
	}

	if e, _ := list.PeekFront(); e == 1 {
		t.Log("Front Passed")
	} else {
		t.Error("PeekFront Failed")
	}

	if e, _ := list.PopBack(); e == 999 {
		t.Log("PeekBack Passed")
	} else {
		t.Error("PeekBack Failed")
	}

	if e, _ := list.PeekBack(); e == 998 {
		t.Log("PeekBack Passed")
	} else {
		t.Error("PeekBack Failed")
	}

	if val, exists := list.Contains(897); exists && val.Value() == 897 {
		t.Log("Contains Passed")
	} else {
		t.Error("Contains Failed")
	}

	if val := list.Remove(765); val == 765 {
		t.Log("Remove Passed")
	} else {
		t.Error("Remove Failed")
	}

	list.Clear()

	if val := list.Remove(9); val == nil {
		t.Log("Remove Passed")
	} else {
		t.Error("Empty Remove Failed")
	}

	if _, err := list.PopBack(); err != nil {
		t.Log("PopBack error detection working")
	} else {
		t.Error("PopBack error detection not working")
	}

	if _, err := list.PopFront(); err != nil {
		t.Log("PopFront error detection working")
	} else {
		t.Error("PopFront error detection not working")
	}
}

func TestDoublyList(t *testing.T) {
	done := make(chan bool)
	list := New()
	var produced, consumed int

	go func() {
		for i := 0; i < capacity*10; i++ {
			list.PushFront(rand.Int())
			produced++
			time.Sleep(time.Microsecond)
		}
		done <- true
	}()

	go func() {
		for i := 0; i < capacity*10; i++ {
			list.PushBack(rand.Int())
			produced++
			time.Sleep(time.Microsecond)
		}
		done <- true
	}()

	go func() {
		for {
			if _, err := list.PopFront(); err != nil {
				break
			}
			time.Sleep(time.Microsecond)
			consumed++
		}
		done <- true
	}()

	go func() {
		for {
			if _, err := list.PopBack(); err != nil {
				break
			}
			time.Sleep(time.Microsecond)
			consumed++
		}
		done <- true
	}()

	for i := 0; i < 4; i++ {
		<-done
	}

	if consumed != produced {
		t.Error("Enqueue not linearizable; test failed")
	} else {
		t.Log("One test passed")
	}
}

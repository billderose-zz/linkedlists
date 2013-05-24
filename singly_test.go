package linked

import (
	"math/rand"
	"testing"
	"time"
)

const capacity = 10000


func TestList(t *testing.T) {
	done := make(chan bool)
	list := new(SinglyLinkedList)
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
		for {
			if _, err := list.Pop(); err != nil {
				break
			}
			time.Sleep(time.Microsecond)
			consumed++
		}
		done <- true
	}()

	for i := 0; i < 2; i++ {
		<-done
	}

	if produced != consumed {
		t.Error("Enqueue not linearizable; test failed")
	} else {
		t.Log("One test passed")
	}
}

func BenchmarkList(b *testing.B) {
	l := new(SinglyLinkedList)
	for j := 0; j < b.N; j++ {
		l.PushFront(j)
		l.Pop()
	}
}

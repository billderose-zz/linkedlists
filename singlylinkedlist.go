package linked

import (
	"errors"
)

var sem = make(chan int, 1)

type SinglyLinkedList struct {
	head   *SingleLink
	length int
}

// Returns reference to new SinglyLinkedList
func NewSinglyLinkedList() *SinglyLinkedList {
	return new(SinglyLinkedList).Clear()
}

// Determine if list is empty
func (ll *SinglyLinkedList) IsEmpty() bool {
	return ll.length == 0
}

// Return the size of the list
func (ll *SinglyLinkedList) Size() int {
	return ll.length
}

// Resets the list
func (ll *SinglyLinkedList) Clear() *SinglyLinkedList {
	ll.head = nil
	ll.length = 0
	return ll
}

// Add interface value to front of list
func (ll *SinglyLinkedList) PushFront(i interface{}) {
	sem <- 1
	ll.put(i)
	<-sem
}

// Remove interface value to front of list
func (ll *SinglyLinkedList) Pop() (interface{}, error) {
	sem <- 1
	val, err := ll.take()
	<-sem
	return val, err
}

// Add interface value to back of list
func (ll *SinglyLinkedList) PushBack(i interface{}) {
	var tail *SingleLink
	for e := ll.head; e != nil; e = e.next {
		tail = e
	}
	link := &SingleLink{i, nil}
	tail.next = link
}

// Return reference to front of list
func (ll *SinglyLinkedList) Front() *SingleLink {
	return ll.head
}

// Private helper method to facilitate pushing
// to the head of the list
func (ll *SinglyLinkedList) put(i interface{}) {
	ll.head = &SingleLink{i, ll.head}
	ll.length++
}

// Private helper method to facilitate poping
// from the head of the list. Returns the value of
// the front elemen
func (ll *SinglyLinkedList) take() (interface{}, error) {
	if !ll.IsEmpty() {
		e := ll.head
		ll.head = ll.head.next
		ll.length--
		return e.val, nil

	}
	return nil, errors.New("List Empty")
}

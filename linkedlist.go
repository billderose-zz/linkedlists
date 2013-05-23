package linked

import (
	"errors"
)

const capacity = 10000

var sem = make(chan int, 1)

type Lister interface {
	Push()
	Pop()
	PushBack()
	Front() interface{}
	Next() interface{}
}

type SingleLink struct {
	val  interface{}
	next *SingleLink
}

type SinglyLinkedList struct {
	head  *SingleLink
	length int
}

func (ll *SinglyLinkedList) Push(i interface{}) {
	sem <- 1
	ll.put(i)
	<-sem
}

func (ll *SinglyLinkedList) Pop() (interface{}, error) {
	sem <- 1
	val, err := ll.take()
	<-sem
	return val, err
}

func (ll *SinglyLinkedList) PushBack(i interface{}) {
	var tail *SingleLink
	for e := ll.head; e != nil; e = e.next {
		tail = e
	}
	link := &SingleLink{i, nil}
	tail.next = link
}

func (ll *SinglyLinkedList) Front() *SingleLink {
	return ll.head
}

func (l *SingleLink) Next() *SingleLink {
	return l.next
}

func (ll *SinglyLinkedList) put(i interface{}) {
	ll.head = &SingleLink{i, ll.head}
	ll.length++
}

func (ll *SinglyLinkedList) take() (interface{}, error) {
	if !ll.isEmpty() {
		e := ll.head
		ll.head = ll.head.next
		ll.length--
		return e.val, nil

	}
	return nil, errors.New("List Empty")
}

func (ll *SinglyLinkedList) isEmpty() bool {
	return ll.length == 0
}


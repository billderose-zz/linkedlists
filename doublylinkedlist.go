package linked

import (
	"errors"
	"reflect"
)

var frontSem, backSem = make(chan int, 1), make(chan int, 1)

type DoublyLinkedList struct {
	head, tail *DoubleLink
	length     int
}

/**
* Returns reference to new DoublyLinkedList
 */
func New() *DoublyLinkedList {
	return new(DoublyLinkedList).Clear()
}

/**
* Determine if list is empty
 */
func (ll *DoublyLinkedList) IsEmpty() bool {
	return ll.length == 0
}

/**
* Return the size of the list
 */
func (ll *DoublyLinkedList) Size() int {
	return ll.length
}

/**
* Resets the list
 */
func (ll *DoublyLinkedList) Clear() *DoublyLinkedList {
	ll.head = nil
	ll.tail = nil
	ll.length = 0
	return ll
}

/**
* Add element to front of list
* @param i -- value to add to list
 */
func (ll *DoublyLinkedList) PushFront(i interface{}) {
	frontSem <- 1
	ll.putFront(i)
	<-frontSem
}

/**
* Add element to back of list
* @param i -- value to add to list
 */
func (ll *DoublyLinkedList) PushBack(i interface{}) {
	backSem <- 1
	ll.putBack(i)
	<-backSem
}

/**
* Remove element from front of list
 */
func (ll *DoublyLinkedList) PopFront() (interface{}, error) {
	frontSem <- 1
	val, err := ll.takeFront()
	<-frontSem
	return val, err
}

/**
* Remove element from back of list
 */
func (ll *DoublyLinkedList) PopBack() (interface{}, error) {
	backSem <- 1
	val, err := ll.takeBack()
	<-backSem
	return val, err
}

/**
* Get, but do not remove, the first element of the list
 */
func (ll *DoublyLinkedList) PeekFront() (interface{}, error) {
	frontSem <- 1
	val, err := ll.takeFront()
	ll.putFront(val)
	<-frontSem
	return val, err
}

/**
* Get, but do not remove, the last element of the list
 */
func (ll *DoublyLinkedList) PeekBack() (interface{}, error) {
	frontSem <- 1
	val, err := ll.takeBack()
	ll.putBack(val)
	<-frontSem
	return val, err
}

/**
* Removes the desired interface value
 */
func (ll *DoublyLinkedList) Remove(i interface{}) interface{} {
	if elm, exists := ll.Contains(i); exists {
		next := elm.Next()
		prev := elm.Prev()
		next.SetPrev(prev)
		prev.SetNext(next)
		elm.SetNext(nil)
		elm.SetPrev(nil)
		return elm.val
	}
	return nil
}

/**
* Return reference to front of  list
 */
func (ll *DoublyLinkedList) Front() *DoubleLink {
	if ll != nil {
		return ll.head
	}
	return nil
}

/**
* Return reference to back of list
 */
func (ll *DoublyLinkedList) Back() *DoubleLink {
	if ll != nil {
		return ll.tail
	}
	return nil
}

/**
* Determine if element is contained in list
 */
func (ll *DoublyLinkedList) Contains(i interface{}) (*DoubleLink, bool) {
	if i != nil {
		for e := ll.Front(); e != nil; e = e.Next() {
			if reflect.DeepEqual(i, e.val) {
				return e, true
			}
		}
	}
	return nil, false
}

/**
* Private helper method to facilitate pushing
* to the head of the list
 */
func (ll *DoublyLinkedList) putFront(i interface{}) {
	if !ll.IsEmpty() {
		head := ll.head
		ll.head = &DoubleLink{i, nil, head}
		head.SetPrev(ll.head)
	} else {
		temp := &DoubleLink{i, nil, nil}
		ll.head = temp
		ll.tail = temp
	}
	ll.length++
}

/**
* Private helper method to facilitate poping
* from the head of the list. Returns the value of
* the front element
 */
func (ll *DoublyLinkedList) takeFront() (interface{}, error) {
	if !ll.IsEmpty() {
		head := ll.head
		ll.head = head.Next()

		// clean up memory
		ll.head.SetPrev(nil)
		head.SetNext(nil)
		ll.length--
		return head.val, nil
	}
	return nil, errors.New("List Empty")
}

/**
* Private helper method to facilitate pushing
* to the tail of the list.
 */
func (ll *DoublyLinkedList) putBack(i interface{}) {
	if !ll.IsEmpty() {
		tail := ll.tail
		ll.tail = &DoubleLink{i, ll.tail, nil}
		tail.SetNext(ll.tail)
		ll.length++
	} else {
		ll.putFront(i)
	}
}

/**
* Private helper method to facilitate poping
* from the tail of the list. Returns the value of the
* back of the list
 */
func (ll *DoublyLinkedList) takeBack() (interface{}, error) {
	if !ll.IsEmpty() {
		tail := ll.tail
		ll.tail = tail.Prev()

		// clean up
		ll.tail.SetNext(nil)
		tail.SetPrev(nil)
		ll.length--
		return tail.val, nil
	}
	return nil, errors.New("List Empty")
}

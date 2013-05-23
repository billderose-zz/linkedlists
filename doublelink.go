package linked

import (
	"errors"
)

type DoubleLink struct {
	val        interface{}
	prev, next *DoubleLink
}

/**
* Return the next element in the list. Returns nil
* if caller is nil
 */
func (l *DoubleLink) Next() *DoubleLink {
	if l != nil {
		return l.next
	}
	return nil
}

/**
* Return the prev element in the list. Returns nil
* if caller is nil
 */
func (l *DoubleLink) Prev() *DoubleLink {
	if l != nil {
		return l.prev
	}
	return nil
}

/**
* Return the value associated with DoubleLink
 */
func (l *DoubleLink) Value() interface{} {
	if l != nil {
		return l.val
	}
	return nil
}

/**
* Set the value associated with DoubleLink
 */
func (l *DoubleLink) SetValue(i interface{}) error {
	if l != nil {
		l.val = i
		return nil
	}
	return errors.New("Nil valued DoubleLink")
}

/**
* Return the next pointer associated with DoubleLink
 */
func (l *DoubleLink) SetNext(next *DoubleLink) error {
	if l != nil {
		l.next = next
		return nil
	}
	return errors.New("Nil valued DoubleLink")
}

/**
* Return the next pointer associated with DoubleLink
 */
func (l *DoubleLink) SetPrev(next *DoubleLink) error {
	if l != nil {
		l.prev = next
		return nil
	}
	return errors.New("Nil valued DoubleLink")
}
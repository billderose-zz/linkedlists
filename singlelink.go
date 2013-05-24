package linked

import "errors"

type SingleLink struct {
	val  interface{}
	next *SingleLink
}

// Return reference to next link in list
func (l *SingleLink) Next() *SingleLink {
	return l.next
}

// Return value held by link
func (l *SingleLink) Value() interface{} {
	return l.val
}

// Set the value associated with DoubleLink
func (l *SingleLink) SetValue(i interface{}) error {
	if l != nil {
		l.val = i
		return nil
	}
	return errors.New("Nil valued DoubleLink")
}

// Set the next pointer associated with DoubleLink
func (l *SingleLink) SetNext(next *SingleLink) error {
	if l != nil {
		l.next = next
		return nil
	}
	return errors.New("Nil valued DoubleLink")
}

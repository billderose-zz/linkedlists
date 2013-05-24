package linked

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

package linked

type SingleLink struct {
	val  interface{}
	next *SingleLink
}

func (l *SingleLink) Next() *SingleLink {
	return l.next
}

func (l *SingleLink) Value() interface{} {
	return l.val
}

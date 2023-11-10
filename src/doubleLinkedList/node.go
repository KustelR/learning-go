package doublelinkedlist

type node struct {
	value any
	next  *node
	prev  *node
}

func (n *node) delete() *node {
	if n.next != nil && n.prev != nil {
		n.next.prev = n.prev
		n.prev.next = n.next
	} else if n.next != nil {
		n.next.prev = nil
	} else {
		n.prev.next = nil
	}

	return n
}

func (n *node) swap(n2 *node) {
	oldNnext := n.next
	oldNprev := n.prev

	n.next = n2.next
	n.prev = n2.prev
	if n2.next != nil {
		n.next.prev = n
	}
	if n2.prev != nil {
		n.prev.next = n
	}
	n2.next = oldNnext
	if oldNnext != nil {
		n2.next.prev = n2
	}
	n2.prev = oldNprev
	if oldNprev != nil {
		n2.prev.next = n2
	}
}

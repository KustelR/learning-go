package linkedList

type node struct {
	value int
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
	nValue := n.value
	n.value = n2.value
	n2.value = nValue
}

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
	/*
		nNext := n.next
		nPrev := n.prev
		n2Next := n2.next
		n2Prev := n2.prev
		if *nNext.prev == *n2Prev {
			n2.next = n2Prev
			n2.prev = nPrev
			n.prev = nNext
			n.next = n2Next
			return
		}
		if nPrev != nil && n2Next != nil && *nPrev == *n2Next {
			n2.prev = n2Next
			n2.next = nNext
			n.next = nPrev
			n.prev = n2Prev
			return
		}
		if n2Next != nil {
			n2Next.prev = n
		}
		if n2Prev != nil {
			n2Prev.next = n
		}
		if nNext != nil {
			nNext.prev = n2
		}
		if nPrev != nil {
			nPrev.next = n2
		}
		n.next = n2Next
		n.prev = n2Prev
		n2.next = nNext
		n2.prev = nPrev
	*/
}

/*
func (n node) String() string {
	return fmt.Sprintf("[DLL node] value: %v, next: %v prev: %v", n.value, n.next, n.prev)
}
*/

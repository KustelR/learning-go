package doublelinkedlist

type DoubleLinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (list *DoubleLinkedList) Prepend(value any) *any {
	list.Length++
	newNode := node{value, list.Head, nil}
	if list.Head == nil {
		list.Head = &newNode
		list.Tail = &newNode
		return &value
	}
	newNode.next = list.Head
	list.Head.prev = &newNode
	list.Head = &newNode

	return &value
}

func (list *DoubleLinkedList) Append(value any) *any {
	list.Length++
	newNode := node{value, nil, nil}
	if list.Tail == nil {
		list.Head = &newNode
		list.Tail = &newNode
		return &value

	}
	newNode.prev = list.Tail
	list.Tail.next = &newNode
	list.Tail = &newNode

	return &value
}

func (list *DoubleLinkedList) Find(value any) int {
	currentEl := list.Head
	for i := 0; currentEl != nil; i++ {
		if currentEl.value != value {
			currentEl = currentEl.next
		} else {
			return i
		}
	}
	return -1
}

func (list *DoubleLinkedList) DeleteByIndex(index int) bool {
	if index == 0 {
		list.Head.delete()
		list.Head = list.Head.next
		list.Length--
		return true
	} else if index == list.Length-1 {
		list.Head.delete()
		list.Head = list.Head.next
		list.Length--
		return true
	}
	element := list.Head
	for i := 0; i < list.Length; i++ {
		if i == index {
			element.delete()
			list.Length--
			return true
		}
		element = element.next
	}
	return false
}

func (list *DoubleLinkedList) Delete(value any) bool {
	isFound := false
	element := list.Head
	i := 0
	for ; i < list.Length; i++ {
		if element.value == value {
			isFound = true
			break
		}
		element = element.next
	}
	if !isFound {
		return false
	}
	if i == 0 {
		list.Head = list.Head.next
		list.Length--
	} else if i == list.Length-1 {
		list.Head = list.Head.next
		list.Length--
	}
	element.delete()
	return true
}

func (list *DoubleLinkedList) Reverse() *DoubleLinkedList {
	if list.Head == nil {
		return list
	}
	prevNode := list.Tail
	for prevNode != nil {
		next := prevNode.next
		prevNode.next = prevNode.prev
		prevNode.prev = next
		prevNode = prevNode.next
	}
	head := list.Head
	list.Head = list.Tail
	list.Tail = head

	return list
}

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

func Create() DoubleLinkedList {
	return DoubleLinkedList{nil, nil, 0}
}

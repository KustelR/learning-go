package doublelinkedlist

type DoubleLinkedList struct {
	Head *node
	Tail *node
}

func (list *DoubleLinkedList) Prepend(value any) *any {
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

func (list *DoubleLinkedList) DeleteByIndex(index int) any {
	return 1
}

func (list *DoubleLinkedList) Delete(value any) int {
	return 1
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

func Create() DoubleLinkedList {
	return DoubleLinkedList{nil, nil}
}

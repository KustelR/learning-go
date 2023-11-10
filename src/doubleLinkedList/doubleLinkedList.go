package doublelinkedlist

type DoubleLinkedList struct {
	values []node
	Head   *node
	Tail   *node
}

func (list *DoubleLinkedList) Prepend(value any) *any {
	newNode := node{value, list.Head, nil}
	newValues := append(make([]node, 0, len(list.values)+1), newNode)
	if list.Head == nil {
		list.Head = &newNode
		list.Tail = &newNode
		list.values = append(list.values, newNode)
		return &value
	}
	newNode.next = list.Head
	list.Head.prev = &newNode
	list.Head = &newNode
	list.values = append(newValues, list.values...)

	return &value
}

func (list *DoubleLinkedList) Append(value any) *any {
	newNode := node{value, nil, nil}
	if list.Tail == nil {
		list.Head = &newNode
		list.Tail = &newNode
		list.values = append(list.values, newNode)
		return &value

	}
	newNode.prev = list.Tail
	list.Tail.next = &newNode
	list.Tail = &newNode

	list.values = append(list.values, newNode)
	return &value
}

func (list *DoubleLinkedList) Reverse() *DoubleLinkedList {
	newValues := make([]node, 0, len(list.values))
	if list.Head == nil {
		return list
	}
	prevNode := list.Tail
	for prevNode != nil {
		next := prevNode.next
		prevNode.next = prevNode.prev
		prevNode.prev = next
		newValues = append(newValues, *prevNode)
		prevNode = prevNode.next
	}
	list.values = newValues
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
	return DoubleLinkedList{make([]node, 0), nil, nil}
}

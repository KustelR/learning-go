package stack

import (
	"fmt"
)

type DoubleLinkedList struct {
	values []node
	head   *node
}

func (list *DoubleLinkedList) push(value any) any {
	newNode := node{value, list.head, nil}
	list.values = append(list.values, newNode)
	if list.head != nil {
		list.head.next = &newNode
	}
	list.head = &newNode

	return value
}

func (list *DoubleLinkedList) pop() any {
	head := list.head
	list.head = list.head.prev
	return head.value
}

type node struct {
	value any
	next  *node
	prev  *node
}

func Test() {
	list := DoubleLinkedList{make([]node, 0), nil}
	fmt.Println(list)
	list.push(1)
	fmt.Println(list)
	fmt.Println(list.pop())
	fmt.Println(list)
	fmt.Println("Alloha")
}

package linkedListInt

import (
	"fmt"
)

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
}

/*
func (n node) String() string {
	return fmt.Sprintf("[DLL node] value: %v, next: %v prev: %v", n.value, n.next, n.prev)
}
*/

type DoubleLinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (list *DoubleLinkedList) Prepend(value int) *int {
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

func (list *DoubleLinkedList) Append(value int) *int {
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

func (list *DoubleLinkedList) IndexOf(value any) int {
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

func (list *DoubleLinkedList) Find(index int) any {
	currentEl := list.Head
	for i := 0; i <= index; i++ {
		if i == index {
			return currentEl.value
		}
	}
	return nil
}

func (list *DoubleLinkedList) findNode(index int) *node {
	currentEl := list.Head
	for i := 0; i <= index; i++ {
		if i == index {
			return currentEl
		}
		currentEl = currentEl.next
	}
	return nil
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

func (list *DoubleLinkedList) swap(n1, n2 *node) {
	n1.swap(n2)
	if n1.next == nil {
		list.Tail = n1
	} else if n2.next == nil {
		list.Tail = n2
	}
	if n1.prev == nil {
		list.Head = n1
	} else if n2.prev == nil {
		list.Head = n2
	}
}

func (list *DoubleLinkedList) Swap(i, j int) {
	n1, n2 := list.findNode(i), list.findNode(j)
	list.swap(n1, n2)
}

func (list *DoubleLinkedList) Sort(fn func(a, b int) (res bool)) {
	left := 0
	right := list.Length - 1

	list.quickSort(fn, left, right)
}

func (list *DoubleLinkedList) quickSort(fn func(a, b int) (res bool), start, end int) {
	if start < end {
		//fmt.Println(list)
		isSorted, p := list.partition(fn, start, end)
		if isSorted {
			return
		}
		list.quickSort(fn, start, p)
		list.quickSort(fn, p+1, end)
	}
}

func (list *DoubleLinkedList) partition(fn func(a, b int) (res bool), start, end int) (bool, int) {
	//fmt.Printf("new partition from %v to %v\n", start, end)
	left := list.findNode(start)
	right := list.findNode(end)
	partition := (end-start)/2 + start
	if (end-start)%2 > 0 {
		partition++
	}
	p := list.findNode(partition)

	l := start
	r := end
	//fmt.Println(left, right, p)
	for l <= r {
		for !fn(left.value, p.value) {
			//fmt.Println("l")
			if l == end {
				return true, r
			}
			left = left.next
			l++
		}
		for fn(right.value, p.value) {
			//fmt.Println("r")
			if r == start {
				return true, r
			}
			right = right.prev
			r--
		}
		//fmt.Println(l, r)
		if l >= r {
			//fmt.Println("break")
			break
		}
		left = left.next
		right = right.prev
		//fmt.Println(l, left, r, right)
		//fmt.Println(list, *list.Head)
		list.swap(left.prev, right.next)
		//fmt.Println(list)
		l++
		r--
	}
	return false, r
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

func (list DoubleLinkedList) String() string {
	values := make([]any, 0, list.Length)
	next := list.Head
	for next != nil {
		//fmt.Println(next)
		values = append(values, next.value)
		next = next.next
	}
	return fmt.Sprintf(" Values: %v Head: %p Tail: %p", values, list.Head, list.Tail)
}

func (list *DoubleLinkedList) Concat(list2 *DoubleLinkedList) *DoubleLinkedList {
	mid := *list2.Head
	newTail := *list2.Tail

	list.Tail.next = &mid
	list.Tail = &newTail
	return list
}

func (list *DoubleLinkedList) Slice(start, end int) *DoubleLinkedList {
	newList := Create()
	currentElement := list.Head
	for i := 0; i < list.Length && i <= end; i++ {
		if i == start {
			newHead := *currentElement
			newHead.prev = nil
			newList.Head = &newHead
		}
		if i == end {
			newTail := *currentElement
			newTail.next = nil
			newList.Tail = &newTail
		}

	}
	return &newList
}

func Create() DoubleLinkedList {
	return DoubleLinkedList{nil, nil, 0}
}

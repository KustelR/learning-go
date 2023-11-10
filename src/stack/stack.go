package stack

type Stack struct {
	Head *node
}

func (stack *Stack) Push(value any) *any {
	newNode := node{value, nil}
	if stack.Head == nil {
		stack.Head = &newNode
		return &value
	}
	newNode.next = stack.Head
	stack.Head = &newNode

	return &value
}

func (stack *Stack) Pop() *any {
	if stack.Head == nil {
		return nil
	}
	value := stack.Head.value
	stack.Head = stack.Head.next
	return &value
}

func (stack *Stack) Peek() *any {
	return &stack.Head.value
}

type node struct {
	value any
	next  *node
}

func Create() Stack {
	return Stack{nil}
}

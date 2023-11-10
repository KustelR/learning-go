package stack

type Stack struct {
	values []*node
	Head   *node
}

func (stack *Stack) Push(value any) *any {
	newNode := node{value, nil}
	if stack.Head == nil {
		stack.Head = &newNode
		stack.values = append(stack.values, &newNode)
		return &value
	}
	newNode.next = stack.Head
	stack.Head = &newNode

	stack.values = append(stack.values, &newNode)
	return &value
}

func (stack *Stack) Pop() *any {
	if stack.Head == nil {
		return nil
	}
	value := stack.Head.value
	stack.values = stack.values[:len(stack.values)-1]
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
	return Stack{make([]*node, 0), nil}
}

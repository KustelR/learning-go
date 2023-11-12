package binaryTree

import "fmt"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

func (n *Node) SetValue(value int) {
	n.value = value
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) Add(direction bool, value int) (*Node, error) {
	newNode := &Node{value: value}
	if direction {
		if n.Left != nil {
			return n.Left, fmt.Errorf("trying to rewrite already existing branch")
		}
		n.Left = newNode
	} else {
		if n.Right != nil {
			return n.Right, fmt.Errorf("trying to rewrite already existing branch")
		}
		n.Right = newNode
	}
	return newNode, nil
}

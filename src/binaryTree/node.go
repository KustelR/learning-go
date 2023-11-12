package binaryTree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
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
		if n.left != nil {
			return n.left, fmt.Errorf("trying to rewrite already existing branch")
		}
		n.left = newNode
	} else {
		if n.right != nil {
			return n.right, fmt.Errorf("trying to rewrite already existing branch")
		}
		n.right = newNode
	}
	return newNode, nil
}

package binaryTree

import "fmt"

type BinaryTree struct {
	Root    *Node
	Current *Node
}

func (tree *BinaryTree) ResetCurrent() *Node {
	tree.Current = tree.Root
	return tree.Current
}

func (tree *BinaryTree) Move(direction bool) (*Node, error) {
	prevNode := tree.Current
	if direction {
		tree.Current = tree.Current.left
	} else {
		tree.Current = tree.Current.right
	}
	if tree.Current == nil {
		tree.Current = prevNode
		return nil, fmt.Errorf("trying to move to a nil node")
	}
	return tree.Current, nil
}

func Create(rootValue int) BinaryTree {
	rootNode := &Node{value: rootValue}
	return BinaryTree{Root: rootNode, Current: rootNode}
}

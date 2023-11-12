package binaryTree

import "fmt"

type SearchTree struct {
	tree BinaryTree
}

func CreateSearchTree(rootValue int) SearchTree {
	return SearchTree{Create(rootValue)}
}

func (tree SearchTree) GetRoot() *Node {
	return tree.tree.Root
}

func (tree SearchTree) Add(value int) *Node {
	if value < tree.tree.Current.value {
		_, err := tree.tree.Move(true)
		if err != nil {
			tree.tree.Current.Left = &Node{value: value}
			return tree.tree.Current.Left
		}
		return tree.Add(value)
	} else {
		_, err := tree.tree.Move(false)
		if err != nil {
			tree.tree.Current.Right = &Node{value: value}
			return tree.tree.Current.Right
		}
		return tree.Add(value)
	}
}

func (tree SearchTree) Find(value int) (string, error) {
	result := ""
	notFoundError := fmt.Errorf("value is not present in tree")
	n := tree.GetRoot()

	for i := 0; n.Left != nil || n.Right != nil; i++ {
		if value > n.value {
			if n.Right == nil {
				return result, notFoundError
			}
			n = n.Right
			result += "r"
		} else if value < n.value {
			if n.Right == nil {
				return result, notFoundError
			}
			n = n.Left
			result += "l"
		} else {
			return result, nil
		}
	}

	return result, notFoundError
}

/*
func (tree SearchTree) String() string {
	return tree.tree.String()
}
*/

package binaryTree

import (
	"fmt"
)

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
		tree.Current = tree.Current.Left
	} else {
		tree.Current = tree.Current.Right
	}
	if tree.Current == nil {
		tree.Current = prevNode
		return nil, fmt.Errorf("trying to move to a nil node")
	}
	return tree.Current, nil
}

/*
	! This is vertical print. Pretty cursed in my opinion

	func (tree BinaryTree) String() string {
		tree.ResetCurrent()
		lines := make([]string, 0)
		lines = drawNode(tree.Current, 0, true, &lines)
		lLen := len(lines)
		for i := 0; i < lLen; i++ {
			pad := lLen - i
			whitespaces := ""
			for j := 0; j < pad; j++ {
				whitespaces += " "
			}
			lines[i] = whitespaces + lines[i]
		}
		return strings.Join(lines, "\n")
	}

	func drawNode(n *Node, depth int, side bool, lines *[]string) []string {
		if len(*lines) <= depth {
			*lines = append(*lines, fmt.Sprintf("%v ", n.value))
			*lines = append(*lines, "/ \\")
			if side {
				(*lines)[depth] += "l"
			} else {
				(*lines)[depth] += "r"
			}
		} else {
			if !side {
				(*lines)[depth] += "r"
				(*lines)[depth+1] += ""
			} else {
				(*lines)[depth] += "l"
			}
			(*lines)[depth] += fmt.Sprintf("  %v ", n.value)
			(*lines)[depth+1] += " / \\"
			if !side {
				(*lines)[depth] += "   "
				(*lines)[depth+1] += "   "
			}
		}
		if n.left != nil {
			*lines = drawNode(n.left, depth+2, true, lines)
		}
		if n.right != nil {
			*lines = drawNode(n.right, depth+2, false, lines)
		}
		return *lines
	}

	func draw(n *Node) string {
		result := fmt.Sprintf("%v ", n.value)
		for n.left != nil {
			n = n.left
			result += fmt.Sprintf("%v ", n.value)
		}
		return result
	}
*/

func Create(rootValue int) BinaryTree {
	rootNode := &Node{value: rootValue}
	return BinaryTree{Root: rootNode, Current: rootNode}
}

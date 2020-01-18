package AVLTree

import (
	"fmt"
	"github.com/1005281342/step-by-step-learn-data-struct/BinarySearchTree"
	"testing"
)

func TestNewAVLNode(t *testing.T) {
	root := NewAVLNode(10)
	node := root.TreeNode
	node = root.Add(12)
	node = root.Add(1)
	node = root.Add(14)
	node = root.Add(2)
	node = root.Add(3)
	node = root.Add(39)
	node = root.Add(16)
	fmt.Println(node.Height)
	BinarySearchTree.InOrder(node)

	node = root.Remove(10)
	fmt.Println(node.Height)
	root.TreeNode = node

	fmt.Println("Is BST:", root.IsBST())
	fmt.Println("Is Balanced:", root.IsBalanced(node))
}

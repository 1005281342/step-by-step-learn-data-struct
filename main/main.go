package main

import (
	"github.com/1005281342/step-by-step-learn-data-struct/BinarySearchTree"
)

func main() {

	// Test BinarySearchTree
	node := BinarySearchTree.NewNode(10)
	node = BinarySearchTree.Add(node, 8)
	node = BinarySearchTree.Add(node, 12)
	node = BinarySearchTree.Add(node, 101)
	//node = BinarySearchTree.Add(node, 1)
	node = BinarySearchTree.Add(node, 9)
	node = BinarySearchTree.Add(node, 1011)

	//fmt.Println(BinarySearchTree.Contains(node, 1))
	//fmt.Println(BinarySearchTree.Contains(node, 11))
	//
	////BinarySearchTree.PreOrder(node)
	////BinarySearchTree.PreOrderNR(node)
	//
	//BinarySearchTree.LevelOrder(node)
	//BinarySearchTree.InOrder(node)

	nNode := BinarySearchTree.RemoveMiniNumNode(node)
	BinarySearchTree.InOrder(nNode)

	nNode = BinarySearchTree.RemoveMaxiNumNode(node)
	BinarySearchTree.InOrder(node)
	//ln := HashMap.NewLinkNode()
	//ln.Value = 10
	//elem := HashMap.LinkNode{20, nil}
	//ln = HashMap.AppendLeft(*ln, elem)
	//fmt.Printf("%+v", elem)
	//fmt.Printf("%+v", ln)

	nNode = BinarySearchTree.Remove(node, 101)
	BinarySearchTree.InOrder(node)
}

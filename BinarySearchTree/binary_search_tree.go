package BinarySearchTree

import (
	"fmt"
	"github.com/1005281342/step-by-step-learn-data-struct/Stack"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
	//cnt int		// 如果需要统计重复插入的值
}

func NewNode(val int) *TreeNode {
	node := new(TreeNode)
	node.value = val
	//node.cnt += 1
	return node
}

// 添加节点
func Add(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewNode(val)
	}

	if val < node.value { // 值小于该节点, 在左边插入
		node.left = Add(node.left, val)
	} else if val > node.value { // 值大于该节点, 在右边插入
		node.right = Add(node.right, val)
	}
	return node
}

// 在树中查找元素
func Contains(node *TreeNode, val int) bool {

	if node == nil {
		return false
	}
	if node.value == val {
		return true
	}
	if val < node.value {
		return Contains(node.left, val)
	}
	//if val > node.value {
	return Contains(node.right, val)
	//}
}

// 前序遍历 -递归版
func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	PreOrder(node.left)
	PreOrder(node.right)
}

// 前序遍历 -非递归
func PreOrderNR(node *TreeNode) {
	if node == nil {
		return
	}
	stk := Stack.NewStack()
	stk.Push(node)
	for stk.Len() > 0 {
		cur := stk.Pop().(*TreeNode)
		fmt.Println(cur.value)
		if cur.right != nil {
			stk.Push(cur.right)
		}
		if cur.left != nil {
			stk.Push(cur.left)
		}
	}
}

// 中序遍历 -递归版
func InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PreOrder(node.left)
	fmt.Println(node.value)
	PreOrder(node.right)
}

// 后序遍历 -递归版
func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PreOrder(node.left)
	PreOrder(node.right)
	fmt.Println(node.value)
}

// 中序遍历	//TODO 该优先队列还需要研究下
//func LevelOrder(node *TreeNode) {
//	if node == nil {
//		return
//	}
//	c := 0
//	pq := PriorityQueue.NewPriorityQueue(100)
//	pq.Push(PriorityQueue.NewItem(node, c))
//	for pq.Len() > 0 {
//		cur := pq.Pop().(*PriorityQueue.Item).Value.(*TreeNode)
//		fmt.Println(cur.value)
//		c += 1
//		if cur.left != nil {
//			pq.Push(PriorityQueue.NewItem(cur.left, c))
//		}
//		if cur.right != nil {
//			pq.Push(PriorityQueue.NewItem(cur.right, c))
//		}
//	}
//}

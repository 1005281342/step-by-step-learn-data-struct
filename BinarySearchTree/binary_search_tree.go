package BinarySearchTree

import (
	"fmt"
	"github.com/1005281342/step-by-step-learn-data-struct/QueueLinkedList"
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
	InOrder(node.left)
	fmt.Println(node.value)
	InOrder(node.right)
}

// 后序遍历 -递归版
func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrder(node.left)
	PostOrder(node.right)
	fmt.Println(node.value)
}

// 中序遍历
func LevelOrder(node *TreeNode) {
	if node == nil {
		return
	}
	c := 1
	q := new(QueueLinkedList.Queue)
	q.Enqueue(node)
	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		cur := v.(*TreeNode)
		fmt.Println(cur.value)
		c += 1
		if cur.left != nil {
			q.Enqueue(cur.left)
		}
		if cur.right != nil {
			q.Enqueue(cur.right)
		}
	}
}

// 寻找树中最小的元素所在节点
func MiniNumNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node
	}
	return MiniNumNode(node.left)
}

// 寻找树中最小元素的值
func MiniNumValue(node *TreeNode) int {
	if node != nil {
		return MiniNumNode(node).value
	}
	return -1 // 该树不存在元素
}

// 从树中移除最小元素节点，并返回新的树
func RemoveMiniNumNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.left != nil { // 如果存在左节点，则继续寻找
		node.left = RemoveMiniNumNode(node.left)
		return node
	}
	// 如果不存在左节点，则当前节点就是最小节点，然后判断是否存在右节点
	if node.right != nil { // 存在右节点时
		rNode := node.right
		node.right = nil
		return rNode
	}
	return nil
}

// 寻找元素中最大的元素所在节点
func MaxiNumNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node
	}
	return MaxiNumNode(node.right)
}

// 寻找树中最大元素的值
func MaxiNumValue(node *TreeNode) int {
	if node != nil {
		return MaxiNumNode(node).value
	}
	return -1
}

// 从树中移除最大元素节点，并返回新的树
func RemoveMaxiNumNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		node.right = RemoveMaxiNumNode(node.right)
		return node
	}
	if node.left != nil {
		lNode := node.left
		node.left = nil
		return lNode
	}
	return nil
}

// 从树中移除元素为val的节点
func Remove(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if val < node.value {
		node.left = Remove(node.left, val)
		return node
	}
	if val > node.value {
		node.right = Remove(node.right, val)
		return node
	}
	// val == node.val
	if node.left == nil { // 其左子树为空时
		rNode := node.right
		node.right = nil
		return rNode
	}
	if node.right == nil { // 其右子树为空时
		lNode := node.left
		node.left = nil
		return lNode
	}
	// 左右子树均不为空时，需要找到当前节点的前驱节点或者后继节点然后进行调整，此处找后继节点
	successorNode := MiniNumNode(node.right) // 找到后继节点
	successorNode.right = RemoveMiniNumNode(node.right)
	successorNode.left = node.left
	node.left = nil
	node.right = nil
	return successorNode

	/* 如果要找前驱节点，然后调整
	precursorNode := MaxiNumNode(node.left)
	precursorNode.left = RemoveMaxiNumNode(node.left)
	precursorNode.right = node.right
	node.left = nil
	node.right = nil
	return precursorNode
	*/
}

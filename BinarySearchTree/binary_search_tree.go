package BinarySearchTree

import (
	"fmt"
	"github.com/1005281342/step-by-step-learn-data-struct/QueueLinkedList"
	"github.com/1005281342/step-by-step-learn-data-struct/Stack"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
	//cnt int		// 如果需要统计重复插入的值
	Height int // 树的高度
}

func NewNode(val int) *TreeNode {
	node := new(TreeNode)
	node.Value = val
	//node.cnt += 1
	return node
}

// 添加节点
func Add(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewNode(val)
	}

	if val < node.Value { // 值小于该节点, 在左边插入
		node.Left = Add(node.Left, val)
	} else if val > node.Value { // 值大于该节点, 在右边插入
		node.Right = Add(node.Right, val)
	}
	return node
}

// 在树中查找元素
func Contains(node *TreeNode, val int) bool {

	if node == nil {
		return false
	}
	if node.Value == val {
		return true
	}
	if val < node.Value {
		return Contains(node.Left, val)
	}
	//if val > node.Value {
	return Contains(node.Right, val)
	//}
}

// 前序遍历 -递归版
func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	PreOrder(node.Left)
	PreOrder(node.Right)
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
		fmt.Println(cur.Value)
		if cur.Right != nil {
			stk.Push(cur.Right)
		}
		if cur.Left != nil {
			stk.Push(cur.Left)
		}
	}
}

func InOrderAns(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	InOrderAns(node.Left, ans)
	*ans = append(*ans, node.Value)
	InOrderAns(node.Right, ans)
}

// 中序遍历 -递归版
func InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Println(node.Value)
	InOrder(node.Right)
}

// 后序遍历 -递归版
func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Value)
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
		fmt.Println(cur.Value)
		c += 1
		if cur.Left != nil {
			q.Enqueue(cur.Left)
		}
		if cur.Right != nil {
			q.Enqueue(cur.Right)
		}
	}
}

// 寻找树中最小的元素所在节点
func MiniNumNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Left == nil {
		return node
	}
	return MiniNumNode(node.Left)
}

// 寻找树中最小元素的值
func MiniNumValue(node *TreeNode) int {
	if node != nil {
		return MiniNumNode(node).Value
	}
	return -1 // 该树不存在元素
}

// 从树中移除最小元素节点，并返回新的树
func RemoveMiniNumNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left != nil { // 如果存在左节点，则继续寻找
		node.Left = RemoveMiniNumNode(node.Left)
		return node
	}
	// 如果不存在左节点，则当前节点就是最小节点，然后判断是否存在右节点
	if node.Right != nil { // 存在右节点时
		rNode := node.Right
		node.Right = nil
		return rNode
	}
	return nil
}

// 寻找元素中最大的元素所在节点
func MaxiNumNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Right == nil {
		return node
	}
	return MaxiNumNode(node.Right)
}

// 寻找树中最大元素的值
func MaxiNumValue(node *TreeNode) int {
	if node != nil {
		return MaxiNumNode(node).Value
	}
	return -1
}

// 从树中移除最大元素节点，并返回新的树
func RemoveMaxiNumNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		node.Right = RemoveMaxiNumNode(node.Right)
		return node
	}
	if node.Left != nil {
		lNode := node.Left
		node.Left = nil
		return lNode
	}
	return nil
}

// 从树中移除元素为val的节点
func Remove(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if val < node.Value {
		node.Left = Remove(node.Left, val)
		return node
	}
	if val > node.Value {
		node.Right = Remove(node.Right, val)
		return node
	}
	// val == node.val
	if node.Left == nil { // 其左子树为空时
		rNode := node.Right
		node.Right = nil
		return rNode
	}
	if node.Right == nil { // 其右子树为空时
		lNode := node.Left
		node.Left = nil
		return lNode
	}
	// 左右子树均不为空时，需要找到当前节点的前驱节点或者后继节点然后进行调整，此处找后继节点
	successorNode := MiniNumNode(node.Right) // 找到后继节点
	successorNode.Right = RemoveMiniNumNode(node.Right)
	successorNode.Left = node.Left
	node.Left = nil
	node.Right = nil
	return successorNode

	/* 如果要找前驱节点，然后调整
	precursorNode := MaxiNumNode(node.Left)
	precursorNode.Left = RemoveMaxiNumNode(node.Left)
	precursorNode.Right = node.Right
	node.Left = nil
	node.Right = nil
	return precursorNode
	*/
}

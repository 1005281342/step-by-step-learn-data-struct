package AVLTree

import (
	"github.com/1005281342/step-by-step-learn-data-struct/BinarySearchTree"
	"github.com/1005281342/step-by-step-learn-data-struct/math"
)

type AVLTree interface {
	IsBST() bool                                          // 判断该二叉树是否为二叉树
	IsBalanced(node *BinarySearchTree.TreeNode) bool      // 判断该二叉树是否为一颗平衡二叉树
	GetHeight(node *BinarySearchTree.TreeNode) int        // 获取节点node的高度
	GetBalanceFactor(node *BinarySearchTree.TreeNode) int // 获取节点node的平衡因子

	MiniNumNode(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode       // 查找最小元素所在节点
	MaxiNumNode(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode       // 查找最大元素所在节点
	ElemNode(node *BinarySearchTree.TreeNode, val int) *BinarySearchTree.TreeNode // 查找元素所在节点

	RightRotate(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode // 右旋
	LeftRotate(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode  // 左旋
	ReBalance(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode   // 再平衡

	Add(node *BinarySearchTree.TreeNode, val int) *BinarySearchTree.TreeNode    // 添加元素
	Remove(node *BinarySearchTree.TreeNode, val int) *BinarySearchTree.TreeNode // 移除元素
}

type AVLNode struct {
	TreeNode *BinarySearchTree.TreeNode
}

func NewAVLNode(val int) *AVLNode {
	return &AVLNode{TreeNode: &BinarySearchTree.TreeNode{Value: val, Height: 1}} // 初始化节点高度为1
}

func (a *AVLNode) GetHeight(node *BinarySearchTree.TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func (a *AVLNode) GetBalanceFactor(node *BinarySearchTree.TreeNode) int {
	if node == nil {
		return 0
	}
	return a.GetHeight(a.TreeNode.Left) - a.GetHeight(a.TreeNode.Right)
}

func (a *AVLNode) IsBalanced(node *BinarySearchTree.TreeNode) bool {
	if node == nil {
		return true
	}
	bf := a.GetBalanceFactor(node)
	if bf > 1 || bf < -1 {
		return false
	}
	return a.IsBalanced(node.Left) && a.IsBalanced(node.Right)
}

func (a *AVLNode) IsBST() bool {
	if a.TreeNode == nil {
		return true
	}
	ans := make([]int, 0)
	BinarySearchTree.InOrderAns(a.TreeNode, &ans)
	for i := 1; i < len(ans); i++ {
		if ans[i-1] > ans[i] {
			return false
		}
	}
	return true
}

func (a *AVLNode) MiniNumNode(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode {
	return BinarySearchTree.MiniNumNode(node)
}

func (a *AVLNode) MaxiNumNode(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode {
	return BinarySearchTree.MaxiNumNode(node)
}

func (a *AVLNode) ElemNode(node *BinarySearchTree.TreeNode, val int) *BinarySearchTree.TreeNode {
	if node == nil {
		return nil
	}
	if val < node.Value { // 在左边
		return a.ElemNode(node.Left, val)
	}
	if val > node.Value {
		return a.ElemNode(node.Right, val) // 右边
	}
	return node // val == node.Value
}

func (a *AVLNode) RightRotate(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode {

	if node == nil {
		return nil
	}
	lNode := node.Left
	rNode := lNode.Right

	// 向右旋转
	lNode.Right = node
	node.Left = rNode

	// 更新Height
	node.Height = math.Max(a.GetHeight(node.Left), a.GetHeight(node.Right)) + 1
	lNode.Height = math.Max(a.GetHeight(lNode.Left), a.GetHeight(lNode.Right)) + 1
	return lNode
}

func (a *AVLNode) LeftRotate(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode {
	if node == nil {
		return nil
	}
	rNode := node.Right
	lNode := rNode.Left

	// 向左旋转
	rNode.Left = node
	node.Right = lNode

	// 更新Height
	node.Height = math.Max(a.GetHeight(node.Left), a.GetHeight(node.Right)) + 1
	rNode.Height = math.Max(a.GetHeight(rNode.Left), a.GetHeight(rNode.Right)) + 1
	return rNode
}

func (a *AVLNode) ReBalance(node *BinarySearchTree.TreeNode) *BinarySearchTree.TreeNode {

	if node == nil {
		return nil
	}

	// 更新高度
	node.Height = math.Max(a.GetHeight(node.Left), a.GetHeight(node.Right)) + 1

	// 计算平衡因子
	bf := a.GetBalanceFactor(node)

	// 平衡维护
	if bf > 1 && a.GetBalanceFactor(node.Left) >= 0 { // LL
		return a.RightRotate(node)
	}
	if bf < -1 && a.GetBalanceFactor(node.Right) <= 0 { // RR
		return a.LeftRotate(node)
	}
	if bf > 1 && a.GetBalanceFactor(node.Left) < 0 { // LR
		node.Left = a.LeftRotate(node.Left)
		return a.RightRotate(node)
	}
	if bf < -1 && a.GetBalanceFactor(node.Right) > 0 { // RL
		node.Right = a.RightRotate(node.Right)
		return a.LeftRotate(node)
	}

	return node
}

func (a *AVLNode) Add(node *BinarySearchTree.TreeNode, val int) *BinarySearchTree.TreeNode {
	node = BinarySearchTree.Add(node, val)
	return a.ReBalance(node)
}

func (a *AVLNode) Remove(node *BinarySearchTree.TreeNode, val int) *BinarySearchTree.TreeNode {
	node = BinarySearchTree.Remove(node, val)
	return a.ReBalance(node)
}

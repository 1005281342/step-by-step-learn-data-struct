package BinarySearchTree

type BST interface {
	Add(val int)
	Contains(val int) bool

	MiniNumNode(node *TreeNode) *TreeNode
	MiniNumValue(node *TreeNode) int
	RemoveMiniNumNode(node *TreeNode) *TreeNode

	MaxiNumNode(node *TreeNode) *TreeNode
	MaxiNumValue(node *TreeNode) int
	RemoveMaxiNumNode(node *TreeNode) *TreeNode

	Remove(node *TreeNode, val int) *TreeNode
}

type BSTNode struct {
	Tn *TreeNode
}

func NewBSTNode(val int) *BSTNode {
	return &BSTNode{Tn: NewNode(val)}
}

func (b *BSTNode) Add(val int) {
	b.Tn = Add(b.Tn, val)
}

func (b *BSTNode) Contains(val int) bool {
	return Contains(b.Tn, val)
}

func (b *BSTNode) MiniNumNode(node *TreeNode) *TreeNode {
	return MaxiNumNode(b.Tn)
}

func (b *BSTNode) MiniNumValue(node *TreeNode) int {
	return MiniNumValue(b.Tn)
}

func (b *BSTNode) RemoveMiniNumNode(node *TreeNode) *TreeNode {
	return RemoveMiniNumNode(b.Tn)
}

func (b *BSTNode) MaxiNumNode(node *TreeNode) *TreeNode {
	return MaxiNumNode(b.Tn)
}

func (b *BSTNode) MaxiNumValue(node *TreeNode) int {
	return MaxiNumValue(b.Tn)
}

func (b *BSTNode) RemoveMaxiNumNode(node *TreeNode) *TreeNode {
	return RemoveMaxiNumNode(b.Tn)
}

func (b *BSTNode) Remove(node *TreeNode, val int) *TreeNode {
	return Remove(b.Tn, val)
}

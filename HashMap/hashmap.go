package HashMap

const DEFAULT_BUCKETS_CNT = 16 // 16 - 1 ->  0b1111

// 计算HashCode
func HashCode(s string) uint32 {
	var mask uint32
	mask = 4294967295 // (1 << 32) - 1
	var h uint32
	for _, c := range s {
		h = (h << 5 & mask) | (h >> 27)
		h += uint32(c)
	}
	return h
}

// 获取index
func Index(code uint32) uint32 {
	return code & (DEFAULT_BUCKETS_CNT - 1)
}

// 链表
type LinkNode struct {
	Value uint32
	Next  *LinkNode
}

func NewLinkNode() *LinkNode {
	return &LinkNode{}
}

// 链表头插法
func AppendLeft(header LinkNode, elem LinkNode) *LinkNode {
	elem.Next = &header
	return &elem
}

type Entry struct {
	Key   string
	Value interface{}
}

type BucketWithSlice struct {
	entries []*Entry
}

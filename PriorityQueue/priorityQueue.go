package PriorityQueue

import "container/heap"

type Item struct {
	Value    interface{}
	Index    int
	Priority int
}

func NewItem(val interface{}, priority int) *Item {
	item := new(Item)
	item.Value = val
	item.Priority = priority
	return item
}

type PriorityQueue []*Item

func NewPriorityQueue(cap int) PriorityQueue {
	return make(PriorityQueue, cap)
}

// 实现接口heap.Interface接口
func (pg PriorityQueue) Len() int {
	return len(pg)
}

func (pg PriorityQueue) Less(i, j int) bool {
	return pg[i].Priority < pg[j].Priority
}

func (pg PriorityQueue) Swap(i, j int) {
	pg[i], pg[j] = pg[j], pg[i]
	pg[i].Index = i
	pg[j].Index = j
}

// add x as element Len()
func (pg *PriorityQueue) Push(x interface{}) {
	n := len(*pg)
	c := cap(*pg)
	if n+1 > c {
		npg := make(PriorityQueue, c*2)
		copy(npg, *pg)
		*pg = npg
	}
	*pg = (*pg)[0 : n+1]
	item := x.(*Item)
	item.Index = n
	(*pg)[n] = item
}

// remove and return element Len() - 1.
func (pg *PriorityQueue) Pop() interface{} {
	n := len(*pg)
	c := cap(*pg)
	if n < (c/2) && c > 25 {
		npg := make(PriorityQueue, n, c/2)
		copy(npg, *pg)
		*pg = npg
	}
	item := (*pg)[n-1]
	item.Index = -1
	*pg = (*pg)[0 : n-1]
	return item
}

func (pg *PriorityQueue) PeekAndShift(max int) (*Item, int) {
	if pg.Len() == 0 {
		return nil, 0
	}

	item := (*pg)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}
	heap.Remove(pg, 0)
	return item, 0
}

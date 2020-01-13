package QueueLinkedList

type Node struct {
	data interface{}
	//data int
	next *Node
}

type Queue struct {
	rear *Node
}

// 元素入队
func (list *Queue) Enqueue(i interface{}) {
	data := &Node{data: i}
	if list.rear != nil {
		data.next = list.rear
	}
	list.rear = data
}

// 元素出队
func (list *Queue) Dequeue() (interface{}, bool) {
	if list.rear == nil {
		return 0, false
	}
	if list.rear.next == nil {
		i := list.rear.data
		list.rear = nil
		return i, true
	}
	current := list.rear
	for {
		if current.next.next == nil {
			i := current.next.data
			current.next = nil
			return i, true
		}
		current = current.next
	}
}

// 获取队首元素
func (list *Queue) Peek() (interface{}, bool) {
	if list.rear == nil {
		return 0, false
	}
	return list.rear.data, true
}

// 获取队列全部元素
func (list *Queue) Get() []interface{} {
	var items []interface{}
	current := list.rear
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

// 队列是否为空
func (list *Queue) IsEmpty() bool {
	return list.rear == nil
}

// 清空队列
func (list *Queue) Empty() {
	list.rear = nil
}

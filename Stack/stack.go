package Stack

import "container/list"

type Stack struct {
	lst *list.List
}

func NewStack() *Stack {
	lst := list.New()
	return &Stack{lst}
}

func (stack *Stack) Push(value interface{}) {
	stack.lst.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.lst.Back()
	if e != nil {
		stack.lst.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.lst.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.lst.Len()
}

func (stack *Stack) Empty() bool {
	return stack.lst.Len() == 0
}

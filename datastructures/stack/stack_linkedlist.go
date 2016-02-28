package stack

import (
	"fmt"
)

type node struct {
  value interface{}
  next *node
}

func (n node) String() string {
  return fmt.Sprintf("[value:%v]", n.value)
}

type Stack struct {
	head *node
}

func (self Stack) String() string {
  ret := "stack :: "
  it := self.head
  for it != nil {
    ret = fmt.Sprintf("%v[%v]", ret, it.value)
    it = it.next
  }
  return ret
}

func NewStack() *Stack {
  return &Stack{head: nil}
}

func (self *Stack) Push(element interface{}) {
	self.head = &node{value: element, next: self.head}
}

func (self *Stack) Pop() interface{} {
  ret := self.head.value
  self.head = self.head.next
	return ret
}

func (self Stack) Top() interface{} {
	return self.head.value
}

func (self Stack) Empty() bool {
	return self.head == nil
}

/*
func main() {
	stack := NewStack()

	fmt.Println(fmt.Sprintf("Empty :: %v == true", stack.Empty()))

	for i := 0; i < 10; i += 1 {
    stack.Push(i)
  }
	fmt.Println(stack)

	fmt.Println(fmt.Sprintf("Top :: %v == 9", stack.Top()))
	fmt.Println(fmt.Sprintf("Pop :: %v == 9", stack.Pop()))
	stack.Pop()
	fmt.Println(fmt.Sprintf("Empty :: %v == false", stack.Empty()))
	stack.Pop()
  for i := 0; i < 10; i += 1 {
    stack.Pop()
  }
	fmt.Println(stack)
}*/

package stack

//import (
//	"fmt"
//)

type StackS struct {
	container []interface{}
}

func NewStackS() *StackS {
  return &StackS{container: make([]interface{}, 0)}
}

func (self *StackS) Push(element interface{}) {
	self.container = append(self.container, element)
}

func (self *StackS) Pop() interface{} {
	ret := self.container[len(self.container)-1]
	self.container = append(self.container[:len(self.container)-1])
	return ret
}

func (self StackS) Top() interface{} {
	return self.container[len(self.container)-1]
}

func (self StackS) Empty() bool {
	return len(self.container) == 0
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
	fmt.Println(stack)
}*/

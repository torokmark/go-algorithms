package queue

import (
  "fmt"
)

type node struct {
  value interface{}
  next *node
}

func (n node) String() string {
  return fmt.Sprintf("[%v]", n.value)
}

type Queue struct {
  head *node
  tail *node
}

func NewQueue() *Queue {
  return &Queue{head: nil, tail: nil}
}

func (self Queue) String() string {
  ret := "queue :: "
  it := self.head
  for it != nil {
    ret = fmt.Sprintf("%v[%v]", ret, it.value)
    it = it.next
  }
  return ret
}

func (self *Queue) enqueue(element interface{}) {
  n := &node{value: element, next: nil}
  if self.head == nil && self.tail == nil {
      self.head = n
      self.tail = n
  } else {
    self.tail.next = n
    self.tail = n
  }
}

func (self *Queue) dequeue() interface{} {

  if self.head == nil {
    fmt.Println("Queue is empty")
  }
  if self.head == self.tail {

  }
  return nil
}

/*
func main() {
  queue := NewQueue()

  for i := 0; i < 10; i += 1 {
    queue.enqueue(i)
  }

  for i := 0; i < 5; i += 1 {
    fmt.Println(fmt.Sprintf("%v", queue.dequeue()))
  }
  fmt.Println(queue)
}*/

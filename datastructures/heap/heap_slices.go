package heap

import (
  "fmt"
)

type Heap struct {
  container []string
}

func NewHeap() *Heap {
  return &Heap{container: make([]string, 0)}
}

func (self *Heap) Insert(element string) {
  self.container = append(self.container, element)
  current := len(self.container)

  for self.container[current] < self.container[parent(current)] {
    self.swap(current, parent(current))
    current = parent(current)
  }
}

func (self *Heap) MinHeap() {
  for idx := len(self.container) / 2; idx >= 1; idx -= 1 {
    self.minHeapify(idx)
  }
}

func (self *Heap) Remove() string {
  //popped := self.container[0]
  self.container[0] = self.container[len(self.container) - 1]
  // move container = append
  return ""
}

func parent(idx int) int {
  return idx / 2
}

func left(idx int) int {
  return 2 * idx
}

func right(idx int) int {
  return 2 * idx + 1
}

func (self *Heap) leaf(idx int) bool {
  if idx >= cap(self.container) / 2 && idx <= cap(self.container) {
    return true
  }
  return false
}

func (self *Heap) swap(fidx, sidx int) {
  self.container[fidx], self.container[sidx] = self.container[sidx], self.container[fidx]
}

func (self *Heap) minHeapify(idx int) {
  if self.leaf(idx) {
    if self.container[idx] > self.container[left(idx)] || self.container[idx] > self.container[right(idx)] {
      if self.container[left(idx)] < self.container[right(idx)] {
        self.swap(idx, left(idx))
        self.minHeapify(left(idx))
      } else {
        self.swap(idx, right(idx))
        self.minHeapify(right(idx))
      }
    }
  }
}

func main() {
  h := NewHeap()
  h.Insert("m")
  h.Insert("c")
  h.Insert("a")
  h.Insert("b")
  fmt.Println(h)
}

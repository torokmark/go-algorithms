package bintree

type Order int

const (
    preorder Order = iota
    inorder Order = iota
    postorder Order = iota
)

type BintreeS struct {
    container []interface{}
}

func NewBintreeS() *BintreeS {
    return &BintreeS{container: make([]interface{}, 0)}
}

func (self *BintreeS) Insert(element interface{}) {
    self.container = append(self.container, element)
}

func (self *BintreeS) Remove(element interface{}) {
    idx := self.Search(element)
    if 0 <= idx && idx <= len(self.container) {
        self.container = append(self.container[:idx], self.container[idx + 1:]...)
    }
}

func (self BintreeS) Get(idx int) interface{} {
    if 0 <= idx && idx < len(self.container) {
        return self.container[idx]
    }
    return nil
}

func (self BintreeS) Children(element interface{}) (interface{}, interface{}) {
    idx := self.Search(element)
    if 0 <= idx && idx < len(self.container) / 2 {
        return self.container[2 * idx + 1], self.container[2 * idx + 2]
    }
    return nil, nil
}

func (self BintreeS) Search(element interface{}) int {
    for i := 0; i < len(self.container); i += 1 {
        if element == self.container[i] {
            return i
        }
    }
    return -1
}

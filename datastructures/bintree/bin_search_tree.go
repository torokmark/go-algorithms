package bintree

import (
  "fmt"
)

type OrderType int

const (
    Pre OrderType = 1 + iota
    In
    Post
)

type Action func(n *Node)

type Node struct {
  value string
  left, right *Node
  parent *Node
}

func NewNode(element string) *Node {
  return &Node{value: element, left: nil, right: nil, parent: nil}
}

func (self Node) String() string {
  return fmt.Sprintf("[%v]", self.value)
}

type Bintree struct {
    root *Node
}

func NewBintree() *Bintree {
    return &Bintree{root: nil}
}

/*
func (self Bintree) String() string {
  ret := "tree :: "
  for i := 0; i < len(self.container); i += 1 {
    ret += fmt.Sprintf("[%v]", self.container[i])
  }
  return ret
}*/

func (self *Bintree) Insert(element string) {
    self.root = self.insert(element, self.root, nil)
}

func (self *Bintree) insert(element string, node *Node, parent *Node) *Node {
  if node == nil {
    node = NewNode(element)
    node.parent = parent
  } else if element < node.value {
    node.left = self.insert(element, node.left, node.parent)
  } else {
    node.right = self.insert(element, node.right, node.parent)
  }
  return node
}

func (self *Bintree) Remove(element string) {
  remove(self.root, element)
}

func remove(node *Node, element string) {
  if element < node.value {
    remove(node.left, element)
  } else if node.value < element {
    remove(node.right, element)
  } else {
    if node.left != nil && node.right != nil {
      successor := findMin(node)
      node.value = successor.value
      remove(successor, successor.value)
    } else if node.left != nil {
      replaceNodeInParent(node, node.left)
    } else if node.right != nil {
      replaceNodeInParent(node, node.right)
    } else {
      replaceNodeInParent(node, nil)
    }
  }
}

func findMin(node *Node) *Node {
  currentNode := node
  for currentNode.left != nil {
    currentNode = currentNode.left
  }
  return currentNode
}

func replaceNodeInParent(parent *Node, node *Node) {
  if parent != nil {
    if parent == parent.left {
      parent.left = node
    } else {
      parent.right = node
    }
  }
  if node != nil {
    node.parent = parent
  }
}

func (self Bintree) Search(element string) *Node {
  return search(self.root, element)
}

func search(node *Node, element string) *Node {
  if node != nil && node.value == element {
    return node
  } else if element < node.value {
    return search(node.left, element)
  } else if node.value <= element {
    return search(node.right, element)
  } else {
    return nil
  }
}


func (self *Bintree) Traverse(order OrderType, action Action) {
  switch order {
  case Pre:
    self.preorderTraverse(self.root, action)
  case In:
    self.inorderTraverse(self.root, action)
  case Post:
    self.postorderTraverse(self.root, action)
  }
}

func (self *Bintree) preorderTraverse(node *Node, action Action) {
  if node != nil {
      action(node)
      self.preorderTraverse(node.left, action)
      self.preorderTraverse(node.right, action)
  }
}

func (self *Bintree) inorderTraverse(node *Node, action Action) {
    if node != nil {
        self.inorderTraverse(node.left, action)
        action(node)
        self.inorderTraverse(node.right, action)
    }
}

func (self *Bintree) postorderTraverse(node *Node, action Action) {
    if node != nil {
        self.postorderTraverse(node.left, action)
        self.postorderTraverse(node.right, action)
        action(node)
    }
}

func (self *Bintree) DepthFirstSearch(element string) *Node {
  return depthFirstSearch(self.root, element)
}

func depthFirstSearch(node *Node, element string) *Node {
  var n *Node
  if node != nil && node.value == element {
    return node
  }
  if node.left != nil {
    n = depthFirstSearch(node.left, element)
    if n != nil && n.value == element {
      return n
    }
  }
  if node.right != nil {
    n = depthFirstSearch(node.right, element)
    if n != nil && n.value == element {
      return n
    }
  }
  return nil
}

func (self *Bintree) BreadthFirstSearch(element string) *Node {
  return breadthFirstSearch(self.root, element)
}

func breadthFirstSearch(node *Node, element string) *Node {
  if node == nil {
    return nil
  }
  queue := make([]*Node, 0)
  queue = append(queue, node)
  for len(queue) > 0 {
    n := queue[0]
    if n.value == element {
      return n
    }
    queue = queue[1:]
    if n.left != nil {
      queue = append(queue, n.left)
    }
    if n.right != nil {
      queue = append(queue, n.right)
    }
  }

  return nil
}

func main() {
  b := NewBintree()
  //fmt.Println(b)

  b.Insert("m")
  b.Insert("b")
  b.Insert("a")
  b.Insert("c")
  b.Insert("o")
  b.Insert("n")
  b.Insert("z")
  //fmt.Println(b.Search("c"))

  //fmt.Println(b)
  //fmt.Println("\nPreorder >> ")
  //b.Traverse(Pre, func (n *Node) { fmt.Print(n) })

  fmt.Println(b.BreadthFirstSearch("m"))
/*
  b.Remove("c")
  fmt.Println(b)

  fmt.Println(b.Search("d"))

  fmt.Println("\nPreorder >> ")
  b.Traverse(Pre, func (n *Node) { fmt.Print(n) })
  fmt.Println("\nInorder >> ")
  b.Traverse(In, func (n *Node) { fmt.Print(n) })
  fmt.Println("\nPostorder >> ")
  b.Traverse(Post, func (n *Node) { fmt.Print(n) })
  */
}

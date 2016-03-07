package main

import (
	"fmt"
	"mygithub/go-algorithms/datastructures/stack"
	"mygithub/go-algorithms/datastructures/queue"
	"mygithub/go-algorithms/datastructures/bintree"
	//"mygithub/go-algorithms/datastructures/heap"

)

func main() {
	//stackTest()
	//queueTest()

	//bintreeTest()
	//bintreesTest()
	//heapTest()
	TestAlgorithm()
}


func stackTest() {
	s := stack.NewStackS()
	s.Push("m")
	fmt.Println(s)
}

func queueTest() {
	q := queue.NewQueue()
	fmt.Println(q)
}

func bintreeTest() {
  b := bintree.NewBintree()

  b.Insert("m")
  b.Insert("b")

  b.Insert("a")
  b.Insert("c")
  b.Insert("o")
  b.Insert("n")
  b.Insert("z")

  b.Remove("c")
  fmt.Println(b)

  fmt.Println(b.Search("d"))

  fmt.Println("\nPreorder >> ")
  b.Traverse(bintree.Pre, func (n *bintree.Node) { fmt.Print(n) })
  fmt.Println("\nInorder >> ")
  b.Traverse(bintree.In, func (n *bintree.Node) { fmt.Print(n) })
  fmt.Println("\nPostorder >> ")
  b.Traverse(bintree.Post, func (n *bintree.Node) { fmt.Print(n) })
}

func bintreesTest() {
	b := bintree.NewBintreeS()
	fmt.Println(b)
	b.Insert("m")
	b.Insert("b")
	b.Insert("a")
	b.Insert("c")
	b.Insert("o")
	b.Insert("n")
	b.Insert("z")

	fmt.Println(b.Children("b"))

	fmt.Println(b)
	b.Remove("c")
	fmt.Println(b)

	fmt.Println(b.Search("a"))
	fmt.Println(b.Get(2))
}

/*
func heapTest() {
	h := heap.NewHeap()

	fmt.Println(h)
}
*/

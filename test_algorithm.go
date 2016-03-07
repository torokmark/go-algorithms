package main

import (
	"fmt"
    "mygithub/go-algorithms/algorithm/sorting"
)


func TestAlgorithm() {
    BubbleSortTest()
	HeapSortTest()
	InsertionSortTest()
	MergeSortTest()
	QuickSortTest()
	SelectionSortTest()
}

func BubbleSortTest() {
    arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
    fmt.Println("Unsorted >> ", arr)
    sorting.Bubble(arr)
    fmt.Println("Bubble sorted >> ", arr)
}

func HeapSortTest() {
	arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
	fmt.Println("Unsorted >> ", arr)
	sorting.HeapSort(arr)
	fmt.Println("Heap sorted >> ", arr)
}

func InsertionSortTest() {
	arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
	fmt.Println("Unsorted >> ", arr)
	sorting.Insertion(arr)
	fmt.Println("Insertion sorted >> ", arr)
}

func MergeSortTest() {
	arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
	fmt.Println("Unsorted >> ", arr)
	fmt.Println("Mergesorted >> ", sorting.MergeSort(arr))
}

func QuickSortTest() {
	arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
	fmt.Println("Unsorted >> ", arr)
	//QuickSort(arr)
	fmt.Println("Quicksorted >> ", sorting.QuickSort(arr))
}

func SelectionSortTest() {
	arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
	fmt.Println("Unsorted >> ", arr)
	sorting.Selection(arr)
	fmt.Println("Selection sorted >> ", arr)
}

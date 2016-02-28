package sorting

import (
  "fmt"
)

func HeapSort(arr []int) {
  start := len(arr)
  end := start
  for i := start; i > -1; i -= 1 {
    heapify(arr, i)
  }
  for i := end - 1; i > 0; i -= 1 {
    arr[0], arr[i] = arr[i], arr[0]
    heapify(i, 0)
  }

}

func heapify(arr []int, currIdx int) {
  left := 2 * currIdx + 1
  right := 2 * (currIdx + 1)
  max := currIdx

  if left < len(arr) && arr[currIdx] < arr[left] {
    max = left
  }
  if right < len(arr) && arr[max] < arr[right] {
    max = right
  }
  if max != currIdx {
    //currIdx, max = max, currIdx
    heapify(len(arr), currIdx)
  }
}

func main() {
  arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
  fmt.Println("Unsorted >> ", arr)
  HeapSort(arr)
  fmt.Println("Heap sorted >> ", arr)
}

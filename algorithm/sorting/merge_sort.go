package sorting

import (
  "fmt"
)

func MergeSort(arr []int) []int {
  ret := []int{}
  midIdx := len(arr) / 2

  if len(arr) < 2 {
    return arr
  }

  left := MergeSort(arr[:midIdx])
  right := MergeSort(arr[midIdx:])

  i, j := 0, 0
  for i < len(left) && j < len(right) {
    if left[i] < right[j] {
      ret = append(ret, left[i])
      i += 1
    } else {
      ret = append(ret, right[j])
      j += 1
    }
  }

  ret = append(ret, left[i:]...)
  ret = append(ret, right[j:]...)
  return ret
}

func main() {
  arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
  fmt.Println("Unsorted >> ", arr)
  fmt.Println("Mergesorted >> ", MergeSort(arr))
}

package sorting

import (
    "fmt"
)

func MergeSort(arr []int) []int {
  ret := []int{}

  if len(arr) < 2 {
    return arr
  }

  midIdx := len(arr) / 2
  left := MergeSort(arr[:midIdx])
  right := MergeSort(arr[midIdx:])
  fmt.Println(left)
  fmt.Println(right)


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
  fmt.Println(ret)
  ret = append(ret, left[i:]...)
  ret = append(ret, right[j:]...)
  fmt.Println(ret)
  return ret
}

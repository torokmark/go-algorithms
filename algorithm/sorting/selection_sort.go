package sorting

import (
  "fmt"
)

func Selection(arr []int) {
  len := len(arr)

  for i := len - 1; i > 0; i -= 1 {
    maxIdx := 0
    for idx := 1; idx < i + 1; idx += 1 {
      if arr[idx] > arr[maxIdx] {
        maxIdx = idx
      }
    }

    arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
  }
}

func main() {
  arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
  fmt.Println("Unsorted >> ", arr)
  Selection(arr)
  fmt.Println("Selection sorted >> ", arr)
}

package sorting

import (
  "fmt"
)

func Bubble(arr []int) {
  len := len(arr)

  for i := 0; i < len; i += 1 {
    for j := 1; j < len - i; j += 1 {
      if arr[j - 1] > arr[j] {
        arr[j - 1], arr[j] = swap(arr[j - 1], arr[j])
      }
    }
  }
}

func swap(a, b int) (int, int) {
  return b, a
}

func main() {
  arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
  fmt.Println("Unsorted >> ", arr)
  Bubble(arr)
  fmt.Println("Bubble sorted >> ", arr)
}

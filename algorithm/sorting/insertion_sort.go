package sorting

import (
  "fmt"
)

func Insertion(arr []int) {
  len := len(arr)

  for i := 1; i < len; i += 1 {
    current := arr[i]
    pos := i

    for pos > 0 && arr[pos - 1] > current {
      arr[pos] = arr[pos - 1]
      pos = pos - 1
    }

    arr[pos] = current
  }
}

func main() {
  arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
  fmt.Println("Unsorted >> ", arr)
  Insertion(arr)
  fmt.Println("Insertion sorted >> ", arr)
}

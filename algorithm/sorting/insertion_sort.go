package sorting

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

package sorting

func Bubble(arr []int) {
  len := len(arr)

  for i := 0; i < len; i += 1 {
    for j := 1; j < len - i; j += 1 {
      if arr[j - 1] > arr[j] {
        arr[j - 1], arr[j] = arr[j], arr[j - 1]
      }
    }
  }
}

package sorting

func MergeSort(arr []int) []int {
  ret := []int{}

  if len(arr) < 2 {
    return arr
  }

  midIdx := len(arr) / 2
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

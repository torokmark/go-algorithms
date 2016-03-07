package sorting

func QuickSort(arr []int) []int {
  less := []int{}
  equal := []int{}
  greater := []int{}

  if len(arr) > 1 {
    pivot := arr[0]
    for i := 0; i < len(arr); i += 1 {
      if arr[i] < pivot {
        less = append(less, arr[i])
      }
      if arr[i] == pivot {
        equal = append(equal, arr[i])
      }
      if arr[i] > pivot {
        greater = append(greater, arr[i])
      }
    }
    less = QuickSort(less)
    greater = QuickSort(greater)
    ret := append(less, equal...)
    ret = append(ret, greater...)

    return ret
  } else {
    return arr
  }
}

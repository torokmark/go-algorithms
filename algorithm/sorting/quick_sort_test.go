package sorting

import (
    "testing"
)

func TestQuickSortWithEmptyArray(t *testing.T) {
    t.Log("Empty list >> expected []")
    arr := make([]int, 0)
    actual := QuickSort(arr)
    if len(actual) != 0 {
        t.Errorf("Expected 0 length, actual length %d", len(actual))
    }
}

func TestQuickSortWithOneLengthArray(t *testing.T) {
    t.Log("One length list >> expected [1]")
    arr := []int{1}
    actual := QuickSort(arr)
    if len(actual) != 1 {
        t.Errorf("Expected 1 length, actual length %d", len(actual))
    }
}

func TestQuickSortWithMultiElementArray(t *testing.T) {
    t.Log("Element is greater than the next one")
    arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
    actual := QuickSort(arr)
    for i := 0; i < len(actual) - 1; i += 1 {
        if actual[i] > actual[i + 1] {
            t.Errorf("The array is not ordered")
        }
    }
}

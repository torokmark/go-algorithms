package sorting

import (
    "testing"
)

func TestSelectionSortWithEmptyArray(t *testing.T) {
    t.Log("Empty list >> expected []")
    arr := make([]int, 0)
    Selection(arr)
    if len(arr) != 0 {
        t.Errorf("Expected 0 length, actual length %d", len(arr))
    }
}

func TestSelectionSortWithOneLengthArray(t *testing.T) {
    t.Log("One length list >> expected [1]")
    arr := []int{1}
    Selection(arr)
    if len(arr) != 1 {
        t.Errorf("Expected 1 length, actual length %d", len(arr))
    }
}

func TestSelectionSortWithMultiElementArray(t *testing.T) {
    t.Log("Element is greater than the next one")
    arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
    Selection(arr)
    for i := 0; i < len(arr) - 1; i += 1 {
        if arr[i] > arr[i + 1] {
            t.Errorf("The array is not ordered")
        }
    }
}

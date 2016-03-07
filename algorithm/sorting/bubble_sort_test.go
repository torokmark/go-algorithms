package sorting

import (
    "testing"
)

func TestBubbleWithEmptyArray(t *testing.T) {
    t.Log("Empty list >> expected []")
    arr := make([]int, 0)
    Bubble(arr)
    if len(arr) != 0 {
        t.Errorf("Expected 0 length, actual length %d", len(arr))
    }
}

func TestBubbleWithOneLengthArray(t *testing.T) {
    t.Log("One length list >> expected [1]")
    arr := []int{1}
    Bubble(arr)
    if len(arr) != 1 {
        t.Errorf("Expected 1 length, actual length %d", len(arr))
    }
}

func TestBubbleWithMultiElementArray(t *testing.T) {
    t.Log("Element is greater than the next one")
    arr := []int{6, 3, 5, 3, 1, 10, 9, 4, 8, 6}
    Bubble(arr)
    for i := 0; i < len(arr) - 1; i += 1 {
        if arr[i] > arr[i + 1] {
            t.Errorf("The array is not ordered")
        }
    }
}

package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{6, 5, 6, 3, 2, 1}
	bubbleSort(array)
	t.Log(array)
}

func TestInsertSort(t *testing.T) {
	array := []int{6, 5, 6, 3, 2, 1}
	insertionSort(array)
	t.Log(array)
}

func TestSelectionSort(t *testing.T) {
	array := []int{6, 5, 6, 3, 2, 1}
	selectionSort(array)
	t.Log(array)
}

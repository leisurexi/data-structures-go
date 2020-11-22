package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{6, 5, 6, 3, 2, 1}
	bubbleSort(array)
	t.Log(array)
}

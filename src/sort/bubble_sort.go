package sort

// 冒泡排序
func bubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		changed := false
		for j := 0; j < len(array)-i-1; j++ {
			temp := array[j]
			if array[j] > array[j+1] {
				array[j] = array[j+1]
				array[j+1] = temp
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}

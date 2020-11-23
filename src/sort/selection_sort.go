package sort

/*
	选择排序
	1.原地排序算法：不需要额外存储空间。
	2.不稳定的排序算法：每次需要找出未排序中元素的最小值，并和前面的元素交换位置；
	比如 5,8,5,2,9 这样一组数据，当第一次找到2位最小元素时，与第一个5交换位置，、
	那第一个5和中间位置的5位置就变了，所有不是稳定的排序算法。
*/
func selectionSort(array []int) {
	sortedIndex := 0
	for i := 0; i < len(array); i++ {
		minIndex := sortedIndex
		for j := sortedIndex + 1; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		temp := array[i]
		array[sortedIndex] = array[minIndex]
		array[minIndex] = temp
		sortedIndex++
	}
}

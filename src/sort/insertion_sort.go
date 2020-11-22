package sort

// 插入排序
// 1. 遍历总元素个数减1次，因为第一个元素默认是有序的
// 2. 倒序比较，先比较当前元素和它的前一个元素
// 3. 如果前一个元素比当前元素大，就往后挪一位，再往前比较
// 4. 比较到小于当前元素时或者比较完时，代表需要插入了
func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		a := array[i]
		// 当前元素的上一个元素的下标
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > a {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		array[j+1] = a
	}
}

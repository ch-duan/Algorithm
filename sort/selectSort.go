package sort

//SelectSort 选择排序
func SelectSort(array []int) {
	if len(array) < 2 {
		return
	}
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[min] > array[j] {
				// array[i], array[j] = array[j], array[i]
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}

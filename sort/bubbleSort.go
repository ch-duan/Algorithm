package sort

//BubbleSort 冒泡
func BubbleSort(array []int) {
	if len(array) < 2 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		ok := true
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				ok = false
			}
		}
		if ok {
			return
		}

	}
}

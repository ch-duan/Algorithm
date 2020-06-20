package sort

//QuickSort QuickSort
func QuickSort(array []int) {
	if len(array) < 2 {
		return
	}
	mid, i := array[0], 1
	head, tail := 0, len(array)-1
	for head < tail {
		if array[i] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	array[head] = mid
	QuickSort(array[head+1:])
	QuickSort(array[:head])
}

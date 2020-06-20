package sort

//MergeSort 归并排序
func MergeSort(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}
	num := length / 2
	left := MergeSort(array[:num])
	right := MergeSort(array[num:])
	return merSort(left, right)
}

func merSort(left []int, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

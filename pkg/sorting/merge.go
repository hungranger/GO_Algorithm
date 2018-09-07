package sorting

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	mergeArr := []int{}

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(mergeArr, right...)
		}

		if len(right) == 0 {
			return append(mergeArr, left...)
		}

		if left[0] < right[0] {
			mergeArr = append(mergeArr, left[0])
			left = left[1:]
		} else {
			mergeArr = append(mergeArr, right[0])
			right = right[1:]
		}
	}

	return mergeArr
}

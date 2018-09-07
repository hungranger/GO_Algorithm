package sorting

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minValue := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if minValue > arr[j] {
				minValue = arr[j]
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}

package searching

func LinearSearch(arr []int, key int) int {
	for k, v := range arr {
		if v == key {
			return k
		}
	}
	return -1
}

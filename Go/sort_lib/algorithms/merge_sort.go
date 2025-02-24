package algorithms

func MergeSort(numbers []int) {
	mergeSort(numbers, 0, len(numbers)-1)
}

func mergeSort(numbers []int, low, high int) {
	if low < high {
		middle := (low + high) / 2
		mergeSort(numbers, low, middle)
		mergeSort(numbers, middle+1, high)
		mergeInPlace(numbers, low, middle, high)
	}
}

package algorithms

func QuickSort(numbers []int) {
	quickSort(numbers, 0, len(numbers)-1)
}

func quickSort(numbers []int, low, high int) {
	if low < high {
		pivotIndex := partition(numbers, low, high)
		quickSort(numbers, low, pivotIndex-1)
		quickSort(numbers, pivotIndex+1, high)
	}
}

func partition(numbers []int, low, high int) int {
	pivot := numbers[high]
	i := low - 1

	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}

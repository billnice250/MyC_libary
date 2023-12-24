package algorithms

func MergeSort(numbers []int) {
	mergeSort(numbers, 0, len(numbers)-1)
}

func mergeSort(numbers []int, low, high int) {
	if low < high {
		middle := (low + high) / 2
		mergeSort(numbers, low, middle)
		mergeSort(numbers, middle+1, high)
		merge(numbers, low, middle, high)
	}
}

func merge(numbers []int, low, middle, high int) {
	left := make([]int, middle-low+1)
	right := make([]int, high-middle)

	for i := 0; i < len(left); i++ {
		left[i] = numbers[low+i]
	}

	for i := 0; i < len(right); i++ {
		right[i] = numbers[middle+1+i]
	}

	i := 0
	j := 0
	k := low

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			numbers[k] = left[i]
			i++
		} else {
			numbers[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		numbers[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		numbers[k] = right[j]
		j++
		k++
	}
}

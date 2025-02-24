package algorithms

func mergeInPlace(numbers []int, low, middle, high int) {
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

// Corrected merge function to merge two sorted arrays
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

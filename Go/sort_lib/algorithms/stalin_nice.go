package algorithms

import (
	"container/heap"
)

// IntHeap type to implement the heap.Interface for min-heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// author: Bill Nice G. Havugukuri (@billnice250)
// date: 2025-02-23
// StalinNiceSort sorts the array in-place using the Stalin Nice Sort algorithm.
// The algorithm an orginal variant of the Stalin Sort algorithm that filters out the
// unsorted elements and then merges the sorted elements with the unsorted elements.
// The algorithm has a time complexity of O(n log n) and a space complexity of O(n).
// The algorithm is stable and in-place.
// references:
// See: https://en.wikipedia.org/wiki/Stalin_sort
// See: https://en.wikipedia.org/wiki/Heapsort
// See: https://en.wikipedia.org/wiki/In-place_algorithm
// See: https://en.wikipedia.org/wiki/Sorting_algorithm
// See: https://en.wikipedia.org/wiki/Time_complexity
func StalinNiceSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	// First element always stays
	kept := []int{arr[0]}
	last := arr[0]
	h := &IntHeap{}
	heap.Init(h)

	// Stalin sort filtering pass
	for _, num := range arr[1:] {
		if num >= last {
			kept = append(kept, num)
			last = num
		} else {
			heap.Push(h, num)
		}
	}

	// Extract sorted elements from heap
	sortedHeap := make([]int, 0, h.Len())
	for h.Len() > 0 {
		sortedHeap = append(sortedHeap, heap.Pop(h).(int))
	}

	// Merge the two sorted arrays
	count := 0
	i, j := 0, 0
	for i < len(kept) && j < len(sortedHeap) {
		if kept[i] < sortedHeap[j] {
			arr[count] = kept[i]
			i++
		} else {
			arr[count] = sortedHeap[j]
			j++
		}
		count++
	}

	// Append remaining elements
	for i < len(kept) {
		arr[count] = kept[i]
		i++
		count++
	}
	for j < len(sortedHeap) {
		arr[count] = sortedHeap[j]
		j++
		count++
	}

}

func StalinNiceSortInPlace(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Phase 1: Filter and collect kept elements
	kept := make([]int, 0, len(arr))
	kept = append(kept, arr[0])
	last := arr[0]
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range arr[1:] {
		if num >= last {
			kept = append(kept, num)
			last = num
		} else {
			heap.Push(h, num)
		}
	}

	// Combined extraction and merging
	count := 0
	i := 0
	var currentHeapElement int
	hasHeapElement := false

	// Initialize first heap element if available
	if h.Len() > 0 {
		currentHeapElement = heap.Pop(h).(int)
		hasHeapElement = true
	}

	// Merge while both sources have elements
	for i < len(kept) && hasHeapElement {
		if kept[i] <= currentHeapElement {
			arr[count] = kept[i]
			i++
		} else {
			arr[count] = currentHeapElement
			// Get next heap element if available
			if h.Len() > 0 {
				currentHeapElement = heap.Pop(h).(int)
			} else {
				hasHeapElement = false
			}
		}
		count++
	}

	// Flush remaining kept elements
	for i < len(kept) {
		arr[count] = kept[i]
		i++
		count++
	}

	// Flush remaining heap elements
	for hasHeapElement {
		arr[count] = currentHeapElement
		count++
		if h.Len() > 0 {
			currentHeapElement = heap.Pop(h).(int)
		} else {
			hasHeapElement = false
		}
	}
}

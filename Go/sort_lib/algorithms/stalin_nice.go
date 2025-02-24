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

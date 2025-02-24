package main

import (
	"math/rand/v2"
	"sort_lib/algorithms"
	"testing"
)

func BenchmarkStalinNiceSort_sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := rand.Perm(1000)
		algorithms.StalinNiceSort(data)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := rand.Perm(1000)
		algorithms.MergeSort(data)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := rand.Perm(1000)
		algorithms.QuickSort(data)
	}
}

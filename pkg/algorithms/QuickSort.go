package algorithms

import ds "goDSA/pkg/dataStructures"

func QuickSort[T ds.Ordered](A *[]T, l, r int) {
	if l < r {
		i := Partition(A, l, r)
		QuickSort(A, l, i-1)
		QuickSort(A, i+1, r)
	}
}

func Partition[T ds.Ordered](A *[]T, l, r int) int {
	pivot := (*A)[r]
	i := l - 1
	for j := l; j < r; j++ {
		if (*A)[j] <= pivot {
			i++
			(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		}
	}
	(*A)[i+1], (*A)[r] = (*A)[r], (*A)[i+1]
	return i + 1
}

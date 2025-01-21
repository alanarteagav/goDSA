package algorithms

import ds "github.com/alanarteagav/goDSA/pkg/dataStructures"

func MergeSort[T ds.Ordered](A *[]T, l, r int) {
	if l < r {
		m := (l + r) / 2
		MergeSort(A, l, m)
		MergeSort(A, m+1, r)
		Merge(A, l, m, r)
	}
}

func Merge[T ds.Ordered](A *[]T, l, m, r int) {
	lengthL := m - l + 1
	lengthR := r - m
	L := make([]T, lengthL)
	R := make([]T, lengthR)
	for i := range lengthL {
		L[i] = (*A)[l+i]
	}
	for j := range lengthR {
		R[j] = (*A)[m+1+j]
	}
	i, j := 0, 0
	k := l
	for i < lengthL && j < lengthR {
		if L[i] <= R[j] {
			(*A)[k] = L[i]
			i++
		} else {
			(*A)[k] = R[j]
			j++
		}
		k++
	}
	for i < lengthL {
		(*A)[k] = L[i]
		i++
		k++
	}

	for j < lengthR {
		(*A)[k] = R[j]
		j++
		k++
	}
}

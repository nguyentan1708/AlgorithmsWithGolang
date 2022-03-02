package main

func SelectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			Swap(&a[i], &a[min])
		}
	}
}

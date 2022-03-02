package main

func InterchangeSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				Swap(&a[i], &a[j])
			}
		}
	}
}

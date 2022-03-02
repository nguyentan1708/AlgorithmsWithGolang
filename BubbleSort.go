package main

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				Swap(&a[j], &a[j-1])
			}
		}
	}
}

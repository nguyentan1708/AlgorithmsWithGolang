package main

func Swap(a *int, b *int) {
	var c int = 0
	c = *a
	*a = *b
	*b = c
}

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				Swap(&a[j], &a[j-1])
			}
		}
	}
}

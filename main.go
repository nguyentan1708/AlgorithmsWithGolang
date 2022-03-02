package main

import (
	"fmt"
	"math/rand"
)

func CreateArray() []int {
	var a [10]int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	return a[:]
}

func main() {
	new_array := CreateArray()
	fmt.Println("Mảng ban đầu:")
	fmt.Println(new_array)
	BubbleSort(new_array)
	fmt.Println("Thuật toán BubbleSort:")
	fmt.Println(new_array)
	SelectionSort(new_array)
	fmt.Println("Thuật toán SelectionSort:")
	fmt.Println(new_array)
}

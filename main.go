package main

import (
	"fmt"
	"math/rand"
)

func Swap(a *int, b *int) {
	var c int = 0
	c = *a
	*a = *b
	*b = c
}

func CreateArray() []int {
	var a [20]int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	return a[:]
}

func main() {
	array1 := CreateArray()
	fmt.Println("Mảng 1:")
	fmt.Println(array1)
	BubbleSort(array1)
	fmt.Println("Thuật toán BubbleSort:")
	fmt.Println(array1)
	//=====================//
	array2 := CreateArray()
	fmt.Println("Mảng 2:")
	fmt.Println(array2)
	SelectionSort(array2)
	fmt.Println("Thuật toán SelectionSort:")
	fmt.Println(array2)
	//====================//
	array3 := CreateArray()
	fmt.Println("Mảng 3:")
	fmt.Println(array3)
	InterchangeSort(array3)
	fmt.Println("Thuật toán InterchangeSort")
	fmt.Println(array3)
}

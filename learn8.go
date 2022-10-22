package main

import (
	"fmt"
)

func BubbleSort(a []uint32) []uint32 {
	for i := range a {
		//print(i)
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func SelectionSort(a []uint32) []uint32 {
	for i := range a {
		minIdx := i
		for j := i + 1; j < len(a); j++ {
			if a[minIdx] > a[j] {
				minIdx = j
			}
		}

		a[i], a[minIdx] = a[minIdx], a[i]
	}

	return a
}

func InsertionSort(a []uint32 ) []uint32 {
	var n = len(a)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j = j - 1
		}
	}

	return a
}

func main() {
	a := []uint32{2, 4, 10, 1, 3, 9, 6}

	fmt.Println(BubbleSort(a))
	fmt.Println(SelectionSort(a))
	fmt.Println(InsertionSort(a))
}

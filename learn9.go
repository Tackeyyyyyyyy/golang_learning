package main

import "fmt"

func BinarySearch(needle uint32, haystack []uint32) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high{
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		}else{
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func main() {
	a := []uint32{2, 4, 10, 1, 3, 9, 6}

	fmt.Println(BinarySearch(7, a))
}
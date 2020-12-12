package main

import (
	"fmt"
)

var unsorted = []int{1,9,2,8,3,7,4,6,5}

func main() {
	sorted := mergeSort(unsorted)
	fmt.Println(sorted)
}

func mergeSort(unsorted []int) []int {
	if len(unsorted) == 1 {
		return unsorted
	}

	middle := len(unsorted)/2
	left := unsorted[:middle]
	right := unsorted[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	size := len(left) + len(right)
	var i int // left
	var j int // right

	result := make([]int, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(left):
			result[k] = right[j]
			j++
		case j == len(right):
			result[k] = left[i]
			i++
		case left[i] > right[j]:
			result[k] = right[j]
			j++
		case left[i] < right[j]:
			result[k] = left[i]
			i++
		case left[i] == right[j]:
			result[k] = right[j]
			j++
		}
	}

	return result
}

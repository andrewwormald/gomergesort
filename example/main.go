package main

import (
	"fmt"

	mergesort "github.com/andrewwormald/gomergesort"
)

var unsorted = []int{1,9,2,8,3,7,4,6,5}

func main() {
	sorted := mergesort.Do(unsorted)
	fmt.Println(sorted)
}



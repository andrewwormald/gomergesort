package mergesort

func Do(unsorted []int) []int {
	if len(unsorted) == 1 {
		return unsorted
	}

	middle := len(unsorted)/2
	left := unsorted[:middle]
	right := unsorted[middle:]

	return merge(Do(left), Do(right))
}

func merge(left, right []int) []int {
	size := len(left) + len(right)
	var i int // left
	var j int // right

	merged := make([]int, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(left):
			merged[k] = right[j]
			j++
		case j == len(right):
			merged[k] = left[i]
			i++
		case left[i] > right[j]:
			merged[k] = right[j]
			j++
		case left[i] < right[j]:
			merged[k] = left[i]
			i++
		case left[i] == right[j]:
			merged[k] = right[j]
			j++
		}
	}

	return merged
}

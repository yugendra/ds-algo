package main

import "fmt"

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}

	iPivot := right

	for i := left; i < iPivot; i++ {
		if array[i] > array[iPivot] {
			tmp := array[i]
			array[i] = array[iPivot-1]
			array[iPivot-1] = array[iPivot]
			array[iPivot] = tmp
			i--
			iPivot = iPivot - 1
		}
	}
	quickSort(array, left, iPivot-1)
	quickSort(array, iPivot+1, right)
}

func main() {
	array := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

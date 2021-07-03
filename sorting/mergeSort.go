package main

import (
	"fmt"
	"math"
)

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	mid := int(math.Round(float64(len(array) / 2)))
	left := array[:mid]
	right := array[mid:]

	return merge(
		mergeSort(left),
		mergeSort(right),
	)

}

func merge(left, right []int) []int {
	var result []int

	il, ir := 0, 0
	lenL := len(left)
	lenR := len(right)

	for il < lenL && ir < lenR {
		if left[il] <= right[ir] {
			result = append(result, left[il])
			il++
		} else {
			result = append(result, right[ir])
			ir++
		}
	}

	if il < lenL {
		result = append(result, left[il:]...)
	}
	/*for i := il; i < lenL; i++ {
		result = append(result, left[i])
	}*/

	if ir < lenR {
		result = append(result, right[ir:]...)
	}
	/*for i := ir; i < lenR; i++ {
		result = append(result, right[i])
	}*/

	return result
}

func main() {
	array := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	sortedArray := mergeSort(array)
	fmt.Println(sortedArray)
}

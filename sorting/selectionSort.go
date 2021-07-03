package main

import "fmt"

func selectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		iSmall := i
		for j := i + 1; j < len(array); j++ {
			if array[iSmall] > array[j] {
				iSmall = j
			}
		}
		tmp := array[i]
		array[i] = array[iSmall]
		array[iSmall] = tmp
	}
}

func main() {
	array := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	selectionSort(array)
	fmt.Println(array)
}

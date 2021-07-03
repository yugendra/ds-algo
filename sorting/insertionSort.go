package main

import "fmt"

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j > -1 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
}

func main() {
	array := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	insertionSort(array)
	fmt.Println(array)
}

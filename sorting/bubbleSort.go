package main

import "fmt"

func bubbleSOrt(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
}

func main() {
	array := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	bubbleSOrt(array)
	fmt.Println(array)
}

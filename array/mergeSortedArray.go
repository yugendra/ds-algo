package main

import "fmt"

func mergeSortedArray(arr1, arr2 []int) []int {

	//Check inputes
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	var sortedArr []int

	var arr1Index int = 0
	var arr2Index int = 0

	for {
		arr1CurItem := arr1[arr1Index]
		arr2CurItem := arr2[arr2Index]

		if arr1CurItem <= arr2CurItem {
			sortedArr = append(sortedArr, arr1CurItem)
			arr1Index++
		} else {
			sortedArr = append(sortedArr, arr2CurItem)
			arr2Index++
		}

		if arr1Index == len(arr1) {
			sortedArr = append(sortedArr, arr2[arr2Index:]...)
			break
		}

		if arr2Index == len(arr2) {
			sortedArr = append(sortedArr, arr1[arr1Index:]...)
			break
		}
	}

	return sortedArr
}

func main() {
	fmt.Println(mergeSortedArray([]int{0, 3, 4, 31}, []int{4, 6, 30}))
}

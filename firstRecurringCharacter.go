package main

import "fmt"

func firstRecurringCharacter(arr []int) int {
	itemMap := make(map[int]bool)
	for _, item := range arr {
		if _, ok := itemMap[item]; ok {
			return item
		}
		itemMap[item] = true
	}
	return 0
}

func main() {
	fmt.Println(firstRecurringCharacter([]int{2, 5, 1, 2, 3, 5, 1, 2, 4}))
	fmt.Println(firstRecurringCharacter([]int{2, 1, 1, 2, 3, 5, 1, 2, 4}))
	fmt.Println(firstRecurringCharacter([]int{2, 3, 4, 5}))
}


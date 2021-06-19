package main

import "fmt"

func firstRecurringCharacter(input []int) int {
	m := make(map[int]bool)

	for _, i := range input {
		if _, ok := m[i]; !ok {
			m[i] = true
		} else {
			return i
		}
	}
	return 0
}

func main() {
	input := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	frc := firstRecurringCharacter(input)
	fmt.Println(frc)

	input = []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	frc = firstRecurringCharacter(input)
	fmt.Println(frc)

	input = []int{2, 3, 4, 5}
	frc = firstRecurringCharacter(input)
	fmt.Println(frc)
}

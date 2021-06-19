package main

import "fmt"

func findFactorialIterative(number int) int {
	var fact int = number
	for i := number - 1; i >= 1; i-- {
		fact = fact * i
	}

	return fact
}

func findFactorialRecursive(number int) int {
	if number == 1 {
		return 1
	}
	return number * findFactorialRecursive(number-1)
}

func main() {
	fact := findFactorialIterative(5)
	fmt.Println("Factorial Iterative:", fact)

	fact = findFactorialRecursive(5)
	fmt.Println("Factorial Recursive:", fact)
}

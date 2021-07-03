package main

import "fmt"

func fibonacciIterative(n int) int {
	var prev int = 0
	var num int = 1

	if n == 0 {
		return prev
	}

	if n == 1 {
		return num
	}

	for i := 2; i <= n; i++ {
		tmp := num
		num = prev + num
		prev = tmp
	}

	return num
}

func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func main() {
	fmt.Println(fibonacciIterative(400))
	fmt.Println(fibonacciRecursive(400))
}

package main

import "fmt"

var calculations int

var cache map[int]int = make(map[int]int)

func fib(i int) int {
	calculations++
	if v, ok := cache[i]; ok {
		return v
	} else {
		if i < 2 {
			return i
		} else {
			cache[i] = fib(i-1) + fib(i-2)
			return cache[i]
		}
	}
}

func main() {
	fmt.Println(fib(100))
	fmt.Println(calculations)
}

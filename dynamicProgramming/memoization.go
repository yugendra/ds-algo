package main

import "fmt"

func memoizeAddTo80() func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		} else {
			fmt.Println("Long time")
			cache[n] = 80 + n
			return cache[n]
		}
	}
}

func main() {
	memoized := memoizeAddTo80()
	fmt.Println(memoized(5))
	fmt.Println(memoized(5))
	fmt.Println(memoized(5))
	fmt.Println(memoized(6))
}

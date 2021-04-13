package main

import "fmt"

type array struct {
	lenght int
	data   map[int]string
}

func (a *array) get(index int) string {
	return a.data[index]
}

func (a *array) push(item string) {
	a.data[a.lenght] = item
	a.lenght++
}

func (a *array) pop() string {
	lastItem := a.data[a.lenght-1]
	delete(a.data, a.lenght-1)
	a.lenght--
	return lastItem
}

func (a *array) deteItem(index int) {
	for i := index; i < a.lenght; i++ {
		a.data[i] = a.data[i+1]
	}
	delete(a.data, a.lenght-1)
	a.lenght--
}

func main() {
	arr := &array{3, map[int]string{0: "a", 1: "b", 2: "c"}}
	fmt.Println(arr.get(1))

	arr.push("d")
	fmt.Println(arr.pop())

	fmt.Println(arr)
	arr.deteItem(1)
	fmt.Println(arr)
}

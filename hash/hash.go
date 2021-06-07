package main

import "fmt"

const hashLength int = 50

type HashTable struct {
	data [hashLength][][2]string
}

func NewHashTable() *HashTable {
	h := &HashTable{}
	h.data = [hashLength][][2]string{}
	return h
}

func _hash(s string) int {
	h := 0
	for i, r := range s {
		h = (h + int(r)*i) % hashLength
	}
	return h
}

func (h *HashTable) set(key, value string) {
	index := _hash(key)
	h.data[index] = append(h.data[index], [2]string{key, value})
}

func (h *HashTable) get(key string) string {
	index := _hash(key)
	for _, v := range h.data[index] {
		if v[0] == key {
			return v[1]
		}
	}

	return ""
}

func (h *HashTable) keys() []string {
	var keysArray []string
	for _, v := range h.data {
		for _, innerV := range v {
			keysArray = append(keysArray, innerV[0])
		}
	}

	return keysArray
}

func main() {
	h := NewHashTable()
	h.set("grapes", "10000")
	fmt.Println(h.get("grapes"))
	h.set("apples", "9")
	fmt.Println(h.get("apples"))

	fmt.Println(h.data)

	fmt.Println(h.keys())
}

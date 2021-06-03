package main

import "fmt"

type graph struct {
	numberOfNodes int
	adjacentList  map[int][]int
}

func NewGraph() *graph {
	return &graph{adjacentList: make(map[int][]int)}
}

func (g *graph) addVertex(node int) {
	g.adjacentList[node] = []int{}
	g.numberOfNodes++
}

func (g *graph) addEdge(node1, node2 int) {
	g.adjacentList[node1] = append(g.adjacentList[node1], node2)
	g.adjacentList[node2] = append(g.adjacentList[node2], node1)
}

func (g *graph) showConnections() {
	for node, nodeConnections := range g.adjacentList {
		fmt.Printf("%d ---> %v\n", node, nodeConnections)
	}
}

func main() {
	g := NewGraph()
	g.addVertex(0)
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)
	g.addVertex(6)

	g.addEdge(3, 1)
	g.addEdge(3, 4)
	g.addEdge(4, 2)
	g.addEdge(4, 5)
	g.addEdge(1, 2)
	g.addEdge(1, 0)
	g.addEdge(0, 2)
	g.addEdge(6, 5)

	g.showConnections()

}

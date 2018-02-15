package main

import "linked_list"

type Graph struct {
	adjacencyListArray []linked_list.List
	v int
}

func New(v int) *Graph {
	return &Graph{
		adjacencyListArray: make([]linked_list.List, v),
		v: v,
	}
}

func (g *Graph) AddEdge(from int, to int) {
	g.adjacencyListArray[from].AddFirst(to)
	g.adjacencyListArray[to].AddFirst(from)
}

func (g *Graph) To2DArray() [][]int {
	arr := make([][]int, g.v)

	for i, list := range g.adjacencyListArray {
		arr[i] = make([]int, g.v)
		for _, j := range list.ToArray() {
			arr[i][j.(int)] = 1
		}
	}

	return arr
}
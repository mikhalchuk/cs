package main

import (
	"testing"
)

func TestGraph_AddEdge(t *testing.T) {
	v := 3
	graph := New(v)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)

	expected := [][]int{
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 0},
	}

	// make sure that Graph is undirected
	array := graph.To2DArray()
	for i := range expected {
		if expected[i][i] != array[i][i] {
			t.Errorf("Should match")
		}
	}
}
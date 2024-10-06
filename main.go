package main

import (
	"fmt"
	"graph-in-golang/utils"
	GraphList "graph-in-golang/utils/Graph"
)

func main() {

	graph, err := utils.Graph_by_file("exemple.graph")

	if err != nil {
		fmt.Println(err)
	}

	utils.Kruskal(graph)

	heap := utils.CreateHeap(10)

	v := GraphList.Vertice{
		Pai:       0,
		Key:       5,
		Info:      1,
		Neighbors: nil,
	}
	utils.AddHeap(&heap, v)

	fmt.Println(heap.Heap[0])

	v = GraphList.Vertice{
		Pai:       0,
		Key:       4,
		Info:      10,
		Neighbors: nil,
	}
	utils.AddHeap(&heap, v)
	v = GraphList.Vertice{
		Pai:       0,
		Key:       3,
		Info:      13,
		Neighbors: nil,
	}
	utils.AddHeap(&heap, v)
	v = GraphList.Vertice{
		Pai:       0,
		Key:       2,
		Info:      24,
		Neighbors: nil,
	}
	utils.AddHeap(&heap, v)

	v = GraphList.Vertice{
		Pai:       0,
		Key:       6,
		Info:      24,
		Neighbors: nil,
	}
	utils.AddHeap(&heap, v)

	v = GraphList.Vertice{
		Pai:       0,
		Key:       1,
		Info:      -1,
		Neighbors: nil,
	}
	utils.AddHeap(&heap, v)

	fmt.Println(heap.Heap)
	fmt.Println(heap.Length)
}

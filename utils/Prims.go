package utils

import (
	"fmt"
	GraphList "graph-in-golang/utils/Graph"
	"math"
)

func Prim(graph **GraphList.GraphList, r int) {
	len := len((*graph).Vertice)

	i := 0
	for i < len {
		(*graph).Vertice[i].Pai = math.MaxInt
		(*graph).Vertice[i].Key = math.MaxInt
		i++
	}

	(*graph).Vertice[r].Key = 0

	heap := CreateHeap(len)
	inHeap := make([]bool, len)
	i = 0
	for i < len {
		AddHeap(heap, &(*graph).Vertice[i])
		inHeap[i] = true
		i++
	}

	fmt.Println("\nTRUE:")
	for *(*heap).Length != 0 {
		u := Heap_extract_min(&heap)
		inHeap[u.Info] = false

		v := u.Neighbors
		for v != nil {
			vertice := v.Vertice
			if inHeap[vertice.Info] && v.Peso < vertice.Key {
				vertice.Pai = u.Info
				vertice.Key = v.Peso
			}

			v = v.Next
		}
		BuildMinHeap(heap)
	}

	i = 0
	soma := 0.0
	for i < len {
		soma += ((*graph).Vertice[i].Key)
		fmt.Println(i, ": ", (*graph).Vertice[i])
		i++
	}

	fmt.Println("Soma: ", soma)
}

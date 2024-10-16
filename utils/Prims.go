package utils

import (
	"fmt"
	GraphList "graph-in-golang/utils/Graph"
	"math"
)

func Prim(graph **GraphList.GraphList, r int) {
	len := len((*graph).Vertice)

	fmt.Println("\nGrafo:")
	GraphList.Print_graph(*graph)
	fmt.Println()

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

	fmt.Print("\nAdicionados na Heap: ")
	for i < len {
		AddHeap(heap, &(*graph).Vertice[i])
		inHeap[i] = true
		fmt.Print(" ", i, " ")
		i++
	}
	fmt.Println("\n")

	for *(*heap).Length != 0 {
		u := Heap_extract_min(&heap)
		fmt.Println(u.Info, " saiu da heap")

		inHeap[u.Info] = false

		v := u.Neighbors
		for v != nil {
			vertice := v.Vertice
			if inHeap[vertice.Info] && v.Peso < vertice.Key {
				fmt.Println("   Vertice ", vertice.Info, " atualizado Pai e Key")
				vertice.Pai = u.Info
				vertice.Key = v.Peso
			}
			v = v.Next
		}
	}
	fmt.Println("\n")

	i = 0
	soma := 0.0
	for i < len {
		v := (*graph).Vertice[i]
		soma += (v.Key)
		fmt.Println(i, ": ", "Key: ", v.Key, " | Pai: ", v.Pai)
		i++
	}

	fmt.Println("\nSoma de pesos: ", soma)
}

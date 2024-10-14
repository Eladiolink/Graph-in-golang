package GraphList

import (
	"fmt"
	"math"
)

type Node struct {
	Vertice *Vertice
	Peso    float64
	Next    *Node
}

type Vertice struct {
	Pai       int
	Key       float64
	Info      int
	Neighbors *Node
}

type GraphList struct {
	Vertice   []Vertice
	IsDigraph bool
}

func Init_graph(qtd_vertex int) GraphList {
	graph := GraphList{
		Vertice:   make([]Vertice, qtd_vertex),
		IsDigraph: true,
	}

	i := 0

	for i < qtd_vertex {
		graph.Vertice[i] = Vertice{
			Pai:       -1,
			Key:       math.MaxInt,
			Info:      i,
			Neighbors: nil,
		}
		i++
	}

	return graph
}

func Add_edge(graph *GraphList, u int, v int, peso float64) bool {
	if len(graph.Vertice) == 0 {
		fmt.Println("O grafo está vazio.")
		return false
	}

	add_neighbors(graph, u, v, peso)

	if graph.IsDigraph {
		add_neighbors(graph, v, u, peso)
	}

	return true
}

func add_neighbors(graph *GraphList, u int, v int, peso float64) bool {
	newNode := Node{
		Vertice: &graph.Vertice[v],
		Peso:    peso,
		Next:    nil,
	}

	if graph.Vertice[u].Neighbors == nil {
		graph.Vertice[u].Neighbors = &newNode
		return true
	}

	newNode.Next = graph.Vertice[u].Neighbors
	graph.Vertice[u].Neighbors = &newNode

	return true
}

func Print_graph(graph *GraphList) {
	lem := len(graph.Vertice)

	i := 0
	for i < lem {
		fmt.Print(i, " : ")
		aux := graph.Vertice[i].Neighbors

		for aux != nil {
			fmt.Print(aux.Vertice.Info, " -> ")
			aux = aux.Next
		}

		fmt.Print("x\n")
		i++
	}
}

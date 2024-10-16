package utils

import (
	"fmt"
	"sort"
)

func Kruskal(G Graph) []Edge {
	edge := make([]Edge, len(G.Vertices)-1)

	fmt.Println("\nGrafo:")
	fmt.Println("Vertices", G.Vertices)
	fmt.Println("Arestas: ", G.Edges)

	set := make([]*LinkedList, len(G.Vertices))

	fmt.Println("\nCriando DisjointSet:")
	for i, v := range G.Vertices {
		set[i] = CreateDisjointset(v)
		fmt.Print(i, " ")
	}
	fmt.Println()

	edges := make([]Edge, len(G.Edges))
	copy(edges, G.Edges)
	sort.Sort(ByPeso(edges))

	count := 0
	fmt.Println()
	for _, aresta := range edges {
		if FindSet(*set[aresta.u]) != FindSet(*set[aresta.v]) {
			fmt.Println("Fazendo união entre:", aresta.u, "-", aresta.v)
			edge[count] = aresta
			Union(&set[aresta.u], &set[aresta.v])
			count++
		}
	}

	// resumo
	var headers []*Header
	lastHeader := set[0].Header
	headers = append(headers, set[0].Header)

	for i := 1; i < len(edge); i++ {
		if lastHeader == set[i].Header {
			continue
		}
		lastHeader = set[i].Header
		headers = append(headers, set[i].Header)

	}

	fmt.Println()
	for i := 0; i < len(headers); i++ {
		fmt.Print("MST ", i+1, ": ")
		PrintSet(headers[i].Head)
	}

	return edge
}

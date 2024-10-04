package utils

import (
	"fmt"
	"sort"
)

func Kruskal(G Graph) []Edge {
	edge := make([]Edge, len(G.Vertices)-1)

	set := make([]*LinkedList, len(G.Vertices))

	for i, v := range G.Vertices {
		set[i] = CreateDisjointset(v)
	}

	edges := make([]Edge, len(G.Edges))
	copy(edges, G.Edges)
	sort.Sort(ByPeso(edges))

	count := 0
	for _, aresta := range edges {
		if FindSet(*set[aresta.u]) != FindSet(*set[aresta.v]) {
			edge[count] = aresta
			Union(&set[aresta.u], &set[aresta.v])
			count++
		}
	}
	fmt.Println(edge)

	return edge
}

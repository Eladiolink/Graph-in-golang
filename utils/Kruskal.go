package utils

import (
	"fmt"
	"sort"
)

func Kruskal(G Graph) {
	edge := make([]Edge, 0)

	set := make([]*LinkedList, 0)

	for _, v := range G.Vertices {
		set = append(set, CreateDisjointset(v))
	}

	edges := make([]Edge, len(G.Edges))
	copy(edges, G.Edges)
	sort.Sort(ByPeso(edges))

	fmt.Println(edges)

	for _, aresta := range edges {
		if FindSet(*set[aresta.u]) != FindSet(*set[aresta.v]) {
			edge = append(edge, aresta)
			Union(&set[aresta.u], &set[aresta.v])
		}
	}

	fmt.Println(edge)
}

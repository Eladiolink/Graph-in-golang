package main

import (
	"fmt"
	"graph-in-golang/utils"
)

func main() {

	graph, err := utils.Graph_by_file("exemple-2.graph")

	if err != nil {
		fmt.Println(err)
	}

	utils.Kruskal(graph)

	// graphList, err := utils.Graph_list_by_file("exemple.graph")

	if err != nil {
		fmt.Println(err)
	}

	// utils.Prim(&graphList, 3)
}

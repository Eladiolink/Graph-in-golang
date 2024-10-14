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

	graphList, err := utils.Graph_list_by_file("exemple.graph")

	if err != nil {
		fmt.Println(err)
	}

	GraphList.Print_graph(&graphList)
}

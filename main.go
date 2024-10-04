package main

import (
	"fmt"
	"graph-in-golang/utils"
)

func main() {

	graph, err := utils.Graph_by_file("exemple.graph")

	if err != nil {
		fmt.Println(err)
	}

	utils.Kruskal(graph)
}

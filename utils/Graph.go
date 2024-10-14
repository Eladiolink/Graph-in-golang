package utils

import (
	"bufio"
	"fmt"
	GraphList "graph-in-golang/utils/Graph"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	u    int
	v    int
	peso float64
}

type Graph struct {
	Vertices []int
	Edges    []Edge
}

func init_graph(qtd_vertex int) Graph {
	var graph Graph

	for i := 0; i < qtd_vertex; i++ {
		graph.Vertices = append(graph.Vertices, int(i))
	}

	return graph
}

func add_edge(graph *Graph, u int, v int, peso float64) {
	if len(graph.Vertices) == 0 {
		fmt.Println("O grafo está vazio.")
		return
	}

	newEdge := Edge{
		u:    u,
		v:    v,
		peso: peso,
	}

	graph.Edges = append(graph.Edges, newEdge)
}

func Graph_by_file(nomeArquivo string) (Graph, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return Graph{}, err
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	count := 0
	var graph Graph

	for scanner.Scan() {
		linha := scanner.Text()

		if count == 1 {
			qtd, err := strconv.Atoi(linha)

			if err != nil {
				return Graph{}, err
			}

			graph = init_graph(qtd)

		}

		if count > 1 {
			infos := strings.Split(linha, " ")
			u, _ := strconv.Atoi(infos[0])
			v, _ := strconv.Atoi(infos[1])
			w, _ := strconv.ParseFloat(infos[2], 64)

			add_edge(&graph, u, v, w)
		}

		count++
	}

	if err := scanner.Err(); err != nil {
		return Graph{}, err
	}

	return graph, nil
}

func Graph_list_by_file(nomeArquivo string) (*GraphList.GraphList, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return &GraphList.GraphList{}, err
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	count := 0
	var graph GraphList.GraphList

	for scanner.Scan() {
		linha := scanner.Text()

		if count == 1 {
			qtd, err := strconv.Atoi(linha)

			if err != nil {
				return &GraphList.GraphList{}, err
			}

			graph = GraphList.Init_graph(qtd)

		}

		if count > 1 {
			infos := strings.Split(linha, " ")
			u, _ := strconv.Atoi(infos[0])
			v, _ := strconv.Atoi(infos[1])
			w, _ := strconv.ParseFloat(infos[2], 64)

			GraphList.Add_edge(&graph, u, v, w)
		}

		count++
	}

	if err := scanner.Err(); err != nil {
		return &GraphList.GraphList{}, err
	}

	return &graph, nil
}

// Para ordenar um slice de Edges
type ByPeso []Edge

func (a ByPeso) Len() int           { return len(a) }
func (a ByPeso) Less(i, j int) bool { return a[i].peso < a[j].peso }
func (a ByPeso) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

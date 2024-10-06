package GraphList

type Node struct {
	Vertice Vertice
	Next    *Node
}

type Vertice struct {
	Pai       int
	Key       int
	Info      int
	Neighbors *Node
}

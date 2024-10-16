package utils

import (
	"fmt"
	GraphList "graph-in-golang/utils/Graph"
)

type Heap struct {
	Heap   []*GraphList.Vertice
	Length *int
}

func CreateHeap(maxLength int) *Heap {
	initial_length := 0
	heap := Heap{
		Heap:   make([]*GraphList.Vertice, maxLength+1),
		Length: &initial_length,
	}

	return &heap
}

func Parent(i int) int { return i / 2 }

func Left(i int) int { return 2 * i }

func Right(i int) int { return 2*i + 1 }

func getElement(heap Heap, i int) *GraphList.Vertice { return heap.Heap[i] }

func MinHeapfy(heap *Heap, i int) {
	l := Left(i)
	r := Right(i)

	menor := i

	if l < *(*heap).Length && getElement(*heap, l).Key < getElement(*heap, menor).Key {
		menor = l
	}

	if r < *(*heap).Length && getElement(*heap, r).Key < getElement(*heap, menor).Key {
		menor = r
	}

	if menor != i {
		aux := getElement(*heap, i)
		(*heap).Heap[i] = getElement(*heap, menor)
		(*heap).Heap[menor] = aux
		MinHeapfy(heap, menor)
	}
}

func AddHeap(heap *Heap, vertice *GraphList.Vertice) {
	(*heap).Heap[*(*heap).Length+1] = vertice
	*(*heap).Length += 1

	for i := *(*heap).Length / 2; i >= 1; i-- {
		MinHeapfy(heap, i)
	}
}

func Heap_extract_min(heap **Heap) GraphList.Vertice {
	BuildMinHeap(*heap)
	tam_heap := *(*heap).Length
	min := (*heap).Heap[1]
	(*heap).Heap[1] = (*heap).Heap[tam_heap]
	*(*heap).Length--

	MinHeapfy(*heap, 1)
	return *min
}

func BuildMinHeap(h *Heap) {
	for i := *h.Length / 2; i >= 1; i-- {
		MinHeapfy(h, i)
	}

}

func PrintHeap(h *Heap) {
	i := 1

	for i < *h.Length {
		fmt.Print(h.Heap[i].Info, " -> ")
		i++
	}
	fmt.Println("x\n")
}

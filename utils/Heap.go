package utils

import (
	"fmt"
	GraphList "graph-in-golang/utils/Graph"
)

type Heap struct {
	Heap   []GraphList.Vertice
	Length int
}

func CreateHeap(maxLength int) Heap {
	return Heap{
		Heap:   make([]GraphList.Vertice, maxLength),
		Length: 0,
	}
}

func Parent(i int) int { return i / 2 }

func Left(i int) int { return 2 * i }

func Rigth(i int) int { return 2*i + 1 }

func getElement(heap Heap, i int) GraphList.Vertice { return heap.Heap[i] }

func MinHeapfy(heap *Heap, i int) {
	l := Left(i)
	r := Rigth(i)

	maior := i

	if l <= (*heap).Length && getElement((*heap), l).Key < getElement((*heap), i).Key {
		maior = l
	}

	if r <= (*heap).Length && getElement((*heap), r).Key < getElement((*heap), i).Key {
		maior = r
	}

	if maior != i {
		fmt.Println(maior, i, l, r)
		aux := (*heap).Heap[i]
		(*heap).Heap[i] = (*heap).Heap[maior]
		(*heap).Heap[maior] = aux

		MinHeapfy(heap, i)
	}
}

func AddHeap(heap *Heap, vertice GraphList.Vertice) {
	(*heap).Heap[(*heap).Length] = vertice
	(*heap).Length += 1

	for i := (*heap).Length / 2; i >= 1; i-- {
		MinHeapfy(heap, i-1)
	}

}

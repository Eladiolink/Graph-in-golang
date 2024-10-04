package utils

import (
	"fmt"
)

type LinkedList struct {
	Header *Header
	Info   int
	Next   *LinkedList
}

type Header struct {
	Head *LinkedList
	Tail *LinkedList
}

func CreateDisjointset(elm int) *LinkedList {
	newLinkedList := LinkedList{
		Header: nil,
		Info:   elm,
		Next:   nil,
	}

	newHeader := Header{
		Head: &newLinkedList,
		Tail: &newLinkedList,
	}

	newHeader.Head.Header = &newHeader
	return &newLinkedList
}

func FindSet(Set LinkedList) *Header {
	return Set.Header
}

func Union(u, v **LinkedList) {

	fmt.Println((*u).Info, (*v).Info)
	// fmt.Println(*(*u).Header, *(*v).Header)

	(*u).Header.Tail.Next = (*v).Header.Head

	aux := (*v).Header.Head

	for aux.Next != nil {
		aux.Header = (*u).Header
		aux = aux.Next
	}
	aux.Header = (*u).Header

	(*u).Header.Tail = aux

	PrintSet((*u).Header.Head)
	fmt.Println()
}

func PrintSet(u *LinkedList) {
	aux := u

	for aux != nil {
		fmt.Print(aux.Info)
		fmt.Print(" -> ")
		aux = aux.Next
	}

	fmt.Println()
}

func PrintHeader(header Header) {
	fmt.Println(header.Head.Header.Head.Info)
}

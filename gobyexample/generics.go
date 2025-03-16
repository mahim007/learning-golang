package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elements []T
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}

	return elements
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo: ", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}

// generics in Go. taken from the https://go.dev/blog/deconstructing-type-parameters
func Clone1[E any](s []E) []E {
	// whatever the function body is
	return s
}

func Clone5[S ~[]E, E any](s S) {
	// 
	// return 1;
}
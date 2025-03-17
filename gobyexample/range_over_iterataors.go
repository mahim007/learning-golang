package main

import (
	"fmt"
	"iter"
	"slices"
)

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

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {

			fmt.Println("yield(e.val):", yield(e.val))
			fmt.Println("after first yield")

			if !yield(e.val) {
				fmt.Println("e.val:", e.val)
				return
			}

			fmt.Println("after 2nd yield")
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				fmt.Println("a:", a, " b:", b)
				return
			}

			a, b = b, a + b
		}
	}
}

func range_over_iterators_main()  {
	lst := List[int]{}
	lst.Push(5)
	lst.Push(10)
	lst.Push(15)
	lst.Push(20)

	for e := range lst.All() {
		fmt.Println("printing: e:", e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}

		fmt.Println(n)
	}
}
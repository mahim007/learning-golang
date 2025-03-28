package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func methods_main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	rp.width = 20
	rp.height = 10
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
}
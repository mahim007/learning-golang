package main

import "fmt"

const π = 3.1416
const PI = π

const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "radian"
)

func const_and_vars_main() {
	const TwoPi, HalfPi, Unit2 = 2 * π, 0.5 * π, "degree"

	fmt.Println(π)
	println(PI)
	println(No)
	println(Yes)
	println(MaxDegrees)
	println(Unit)
	println(TwoPi)
	println(HalfPi)
	println(Unit2)
}
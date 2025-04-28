package main

func SquareOfSumAndDiff(a int64, b int64) (int64, int64) {
	return (a + b) * (a + b), (a - b) * (a - b)
}

func CompareLower4bits(m, n uint32) (r bool) {
	r = m&0xF > n&0xF
	return
}

var v = VersionString()

func main_functions() {
	println(v)
	x, y := SquareOfSumAndDiff(3, 6)
	println(x, y)
	b := CompareLower4bits(uint32(x), uint32(y))
	println(b)
	doNothing("Go", 1)
}

func VersionString() string {
	return "v1.0"
}

func doNothing(string, int32) {

}
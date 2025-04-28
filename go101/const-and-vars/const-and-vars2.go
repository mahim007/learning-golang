package main

func main() {
	const MaxUnit uint = (1 << 64) - 1
	println(MaxUnit)

	const MaxUnit2 = ^uint(0)
	println(MaxUnit2)

	const MaxUnit3 = int(^uint(0) >> 1)
	println(MaxUnit3)

	const NativeWordBits = 32 << (^uint(0) >> 63)
	println(NativeWordBits)

	const Is64bitOS = ^uint(0)>>63 != 0
	println("is64bit: ", Is64bitOS)
	const Is32bitOS = ^uint(0)>>32 == 0
	println("is32bit:", Is32bitOS)
}
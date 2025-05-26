package main

import "fmt"

func main() {
	var x uint8 = 15
	var y uint8 = 4
	fmt.Printf("%08b\n", x) // 00001011
	fmt.Printf("%08b\n", y)
	// 位清除运算符，运算结果为 x 与 y 的按位与运算的取反
	fmt.Printf("%08b\n", x&^y) // 00001011
}

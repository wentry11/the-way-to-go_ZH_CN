package main

import "fmt"

func main() {

	for i := 0; i < 11; i++ {
		fmt.Println(fblq(i))
	}

	print10To1()

}

func print10To1() {
	printNum(10)
}

func printNum(i int) {
	if i == 1 {
		fmt.Println(i)
	} else {
		printNum(i - 1)
		println(i)
	}
}

func fblq(idx int) (i int, val int) {
	if idx == 0 || idx == 1 {
		return idx, 1
	} else {
		_, v1 := fblq(idx - 1)
		_, v2 := fblq(idx - 2)
		return idx, v1 + v2
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var orig string = "ABC"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		//return
		//os.Exit(1) // 这里先放开注释，代码继续往下走
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	// 可以将atoi进行封装，忽略错误
	fmt.Println(atoi("abc"))

}

func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

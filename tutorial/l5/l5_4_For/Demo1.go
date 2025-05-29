package main

import "fmt"

func main() {

	// 普通for循环，没有小括号
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 单条件for循环，类似于while
	i := 10
	for i > 0 {
		i--
		fmt.Println(i)
	}

	// 无限循环
	i = 100
	for {
		fmt.Println("hello")
		i--
		if i < 0 {
			break
		}
	}

	// for-range 遍历
	str := "hello world"
	for _, s := range str {
		fmt.Printf("%c", s)
		fmt.Println(string(s))
	}

	// break, continue的用法和Java类似
}

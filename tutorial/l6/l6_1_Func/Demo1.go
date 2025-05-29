package main

import "fmt"

func main() {

	var f F3 = myF3 // 需要和F3具有相同的形参，才能被赋值给F3类型的变量
	f(1, 2)
}

// go 不支持函数的重载，即不同的函数，方法名一定不一样
func f1() {}

//func f1(a int){} // 尽管形参不一样，还是会报错“'f1' redeclared in this package”

func f2() {}

// 只声明一个函数，而没有函数体，F3就是一种函数类型。类似于Java中的接口定义
type F3 func(a int, b int) (int, int)

// 需要和F3具有相同的形参，才能被赋值给F3类型的变量
func myF3(a int, b int) (int, int) {
	fmt.Println("this is myF3")
	return a, b
}

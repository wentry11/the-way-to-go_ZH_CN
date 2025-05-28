package main

func main() {

	F3 := func(a int, b int) {

	}
	F3(1, 2)
}

// go 不支持函数的重载，即不同的函数，方法名一定不一样
func f1() {}

//func f1(a int){} // 尽管形参不一样，还是会报错“'f1' redeclared in this package”

func f2() {}

// 只声明一个函数，而没有函数体
type F3 func(a int, b int) (int, int)

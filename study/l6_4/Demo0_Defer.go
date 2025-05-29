package main

import "fmt"

func main() {
	fmt.Println("defer的作用类似于finally，用于在方法执行结束前做一些操作")
	defer1()

	fmt.Println(defer2())

	/*
		常见的操作:
		1. 关闭文件流
		2. 解锁加锁的资源
		3. 打印日志
		4. 关闭数据库连接
	*/

	printLogExample(2)
}

func printLogExample(i int) {

}

func defer2() int {
	fmt.Println("defer2 enter")
	fmt.Println("defer2 out")
	defer fmt.Println("defer2 defer")
	return 1
}

func defer1() {
	fmt.Println("defer1 enter...")
	defer fmt.Println("defer1 defer...")
	fmt.Println("defer1 out...")
}

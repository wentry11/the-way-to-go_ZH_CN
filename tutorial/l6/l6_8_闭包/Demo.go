package main

import "fmt"

func main() {

	/**
	闭包的意思就是，匿名函数可以直接在后面进行调用。
	尽管可以将其赋值给一个变量，该变量是一个函数值
	*/
	closePackageExample()

	// 这里打印的是几呢？
	// 答案是2，因为defer在return之后执行，将ret直接++了，ret是一个变量
	// 因此defer可以做很多事，包括修改返回的error类型，进行包装或转换等操作
	fmt.Println(testDefer())
}

/*
*
试问：这里返回的ret到底是几？
*/
func testDefer() (ret int) {
	defer func() {
		ret++
	}() //立即调用
	return 1 // 返回1，defer里面有++，实际返回的是2
}

func closePackageExample() {

	func(a int) {
		fmt.Printf("这是一个闭包函数，入参 %d \n", a)
	}(1)

}

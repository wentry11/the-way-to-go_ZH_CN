package main

//import "fmt"
//使用别名
import fm "fmt"

/*
4.2.1 包的概念、导入、可见性
可见性规则：大写开头的都可见（相当于public）

4.2.2 函数
1. 每一个可执行程序必须要有一个main函数
2. 不需要显示的分号；, 都由编译器自行完成
3. gofmt会自动格式化所有的go文件，但编写代码需要考虑可读性
4. function的可见性，通过首写字母的大小写确定。如果不需要暴露给外部用，使用小驼峰即可
5. go没有继承的概念，相较而言，类关系更扁平化
6. fmt.println代表控制台标准输出，也可以使用buildin自带的print函数，但仅限于开发调试阶段
7. 函数执行完成，正常退出的输出是”Process finished with the exit code 0“
   如果是非正常输出，则返回非零值："Compilation finished with exit code 1"


4.2.3 注释
1. 多行注释
2. 单行注释 //
3. 所有的 包、 类 、 函数，最好都要有注释

4.2.4 类型
1. var声明的变量，会自动推导类型，并初始化为该类型的零值
2. 类型可以是
- 基本类型：int、float、bool、string
- 结构化的（复合类型）：struct、array、切片（slice）、map、通道（channel）
- 只描述类型行为的：接口（interface）
3. 结构化的类型没有真正的值，使用nil作为默认值。（go中不存在类型继承！！）
4. 函数的返回值，在最后声明 func FunctionName(a typea, b typeb) typeFunc
5.

*/

func main() {
	//fmt.Printf("hello world")

	//	使用别名
	fm.Println("hello world")

	// 函数调用
	thisIsFirstFunction()

	// 使用buildin的print进行调试
	println("hello")

	// 测试非正常退出
	//var v int = 1 / 0
	//fm.Println(v)

	a := 5
	b := int(a)
	fm.Println(b)
}

func thisIsFirstFunction() {
	fm.Println("this is the first custom function")
}

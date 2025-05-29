package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	p := &i1
	fmt.Println(p)

	fmt.Printf("p的类型为：%T \n", p)

	fmt.Println("还可以通过反射获取类型：")
	typeOf := reflect.TypeOf(p)
	fmt.Println("p的类型是：", typeOf)

	fmt.Println("可以在指针类型前面，再加一个*获取对应的值")
	ptr := &i1        // 获取到指针类型
	fmt.Println(i1)   // 打印变量
	fmt.Println(*ptr) // 再通过*，将指针转换成对应地址变量值

	fmt.Println("对于任何一个变量var，这个表达式都是正确的：var == *&var")
	var a = 1
	fmt.Println(a == *&a)

	fmt.Println("通过指针，可以直接修改变量的值，而不用s = newValue")
	s := "good bye"
	var sPtr *string = &s
	*sPtr = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", sPtr)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *sPtr) // prints string
	fmt.Printf("Here is the string s: %s\n", s)      // prints same string

	// 对空指针，进行赋值，会报错
	var nilPtr *int = nil
	fmt.Println(nilPtr)
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x9cc3dc7]
	*/
	//*nilPtr = 1
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x2c2bdcd]
	*/
	//fmt.Println(*nilPtr)

}

package main

import "fmt"

func main() {

	/**
	1. switch可以是任何类型，case中的类型也要一样
	2. 不需要加break，默认终止执行，如果需要继续往下执行，使用fallthrough
	3. case后面可以跟多个
	*/

	//a := 1
	a := 2
	switch a {
	case 1:
		fmt.Println("1")
		//不需要break
	case 2:
		fmt.Println("2")
		//也不需要break
		//如果需要继续往下执行，使用fallthrough
		fallthrough
	default:
		fmt.Println("none hit")
	}

	/*
		switch的第二种形式是不提供初始值，直接在case中写表达式
		类似于if-else if-else
	*/

	switch {
	case a == 1:
		fmt.Println("a==1")
	case a == 2:
		fmt.Println("a==2")
	default:
		fmt.Println("既不等于一也不等于二")

	}
}

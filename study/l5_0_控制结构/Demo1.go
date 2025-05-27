package main

import (
	"fmt"
)

func main() {

	bool := true
	if bool {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//	判断一个字符串是否是空
	str := ""
	if str == "" {
		fmt.Println("是空的")
	}

	if len(str) == 0 {
		fmt.Println("也是空的")
	}

	fmt.Println(Myabs(1))
	fmt.Println(Myabs(0))
	fmt.Println(Myabs(-1))

	//	if还可以带一个初始化语句，用分号分开
	if v := 1; v > 0 {
		fmt.Println("v 大于 0 ")
	}

	// 也可以从一个函数获取值
	if v := aFun(); v > 0 {
		fmt.Println("函数返回的 大于 0")
	}

}

func aFun() int {
	return 1
}

func Myabs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

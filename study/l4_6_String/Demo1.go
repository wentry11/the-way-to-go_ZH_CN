package main

import "fmt"

func main() {
	fmt.Println("============================================")
	fmt.Println("--1、通过反引号应用的字符串，可以将转移字符原样输出")
	fmt.Println(`原样输出 \n`)
	fmt.Println("转移输出  \n")

	fmt.Println("============================================")
	fmt.Println("--2、通过len获取字符串的长度")
	fmt.Println("length is:", len("hello world"))

	fmt.Println("--3、获取第n位的字符，直接使用中括号，获取出来的是byte")
	fmt.Println("hello world"[1])

}

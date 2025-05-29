package main

import (
	"fmt"
	"strings"
)

func main() {

	// 回调，简单说就是参数里面传的是函数
	callback(1, f1)
	callback(4, f1)

	// 最典型的回调是 strings.IndexFunc
	/**
	该函数 IndexFunc 用于查找字符串 s 中第一个满足函数 f 的 Unicode 码点（字符）的索引位置，若没有符合条件的字符，则返回 -1。
	简要逻辑如下：
	遍历字符串中的每个字符
	对每个字符调用函数 f 判断是否满足条件
	返回第一个满足条件的字符的位置
	如果都没有满足条件的，返回 -1
	*/
	indexFunc := strings.IndexFunc("hello", func(r rune) bool {
		return r == 'h'
	})
	fmt.Println(indexFunc)

	// strings.map同样是一个例子
	/**
	其实就类似于Java中的各种 function、supplier、consumer等，传入一段逻辑作为回调而已
	*/
	s := strings.Map(func(r rune) rune {
		if r > 255 {
			return '?'
		}
		return r
	}, "hello中国")
	fmt.Println(s)
}

func f1(a int, b int) {
	fmt.Printf("%d + %d = %d \n", a, b, a+b)
}

func callback(c int, f func(a int, b int)) {
	f(c, 2) //这里只有一个入参，2是写死的
}

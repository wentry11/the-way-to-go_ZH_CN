package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*
这段 Go 代码的功能如下：
使用 rand.Int() 生成 10 个伪随机整数并输出；
使用 rand.Intn(8) 生成 5 个小于 8 的随机整数并输出；
使用当前时间的纳秒作为种子，设置随机数生成器；
生成 10 个 [0,100) 范围内的随机浮点数并格式化输出。
*/
func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}

	fmt.Println()

	fmt.Printf("%3.2f\n", 1.2)
	fmt.Printf("%3.2f\n", 11.2)
	fmt.Printf("%3.2f\n", 111.2)
	fmt.Printf("%3.2f\n", 1111.2)
}

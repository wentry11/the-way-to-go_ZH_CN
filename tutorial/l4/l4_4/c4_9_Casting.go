package main

import "fmt"

func main() {
	var a int16 = 32 // 定义16位的int
	var b int32      // 定义32位的int
	// 16位的需要显示转换为32位的
	b = int32(a)

	fmt.Printf("16bit int is %d \n", a)
	fmt.Printf("32bit int is %d \n", b)

}

package main

import "fmt"

func main() {

	min := returnMin([]int{3, 2, 3, 1, -1, 5, 6, 7, 9})
	fmt.Println(min)

	// 打印所有的入参的类型
	printTypeAndVar(1, "2", 3.3, nil)

	// printf示例代码
	printFExample()
}

func returnMin(ints []int) int {
	res := ints[0]
	for _, i := range ints {
		if i < res {
			res = i
		}
	}
	return res
}

func printFExample() {
	a := 1

	fmt.Printf("value: %v, type: %T\n", 42, 42) // value: 42, type: int
	fmt.Printf("binary: %b\n", 42)              // binary: 101010
	fmt.Printf("hex: %x\n", 255)                // hex: ff
	fmt.Printf("float: %f\n", 3.1415)           // float: 3.141500
	fmt.Printf("pointer: %p\n", &a)
}

// 不同类型的入参，使用interface接受，通过type进行类型判断
func printTypeAndVar(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println("int:", arg)
		case string:
			fmt.Println("string:", arg)
		case float32:
			fmt.Println("float32:", arg)
		case float64:
			fmt.Println("float64", arg)
		default:
			fmt.Printf("other type %T: %v\n", arg, arg)
		}
	}
}

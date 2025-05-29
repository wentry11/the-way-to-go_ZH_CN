package main

import "fmt"

func main() {

	// 标签
	//labelTest()

	// 使用label 和 goto模拟循环
	useGotoAndLabelSimulateLoop()

	// goto 和 label定义之间，不能出现别的变量声明语句，否则会编译报错
	statementShouldNotExistBetweenGotoAndLabel()
}

func statementShouldNotExistBetweenGotoAndLabel() {
	//    goto LABEL2 // compile error
	//    i := 2
	//LABEL2:
	//	fmt.Println(i)
	//
}

func useGotoAndLabelSimulateLoop() {
	i := 1
HERE:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

/*
*
这是一段死循环代码
*/
func labelTest() {
Label1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				goto Label2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

Label2:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				goto Label1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

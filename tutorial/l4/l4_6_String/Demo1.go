package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

	fmt.Println("--4. 是否有某个前缀，使用strings.HasPrefix")
	fmt.Println("--5. 是否有某个后缀，使用strings.HasSuffix")
	fmt.Println(strings.HasPrefix("ss", "s"))
	fmt.Println(strings.HasSuffix("ss", "s"))

	fmt.Println("--6. 获取字符串的位置 ")
	var str string = "Hi, I'm Marc, Hi."

	fmt.Println(str)
	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	fmt.Println("--7. 字符串的替换，后面的数字代表要替换几个，如果是-1，则替换所有的")
	replace := strings.Replace("aaabbbccc", "a", "d", -1)
	fmt.Println(replace)
	fmt.Println("--8. 字符串的次数统计，此次数为（非重叠次数）")
	count := strings.Count("aaaabbbccc", `aa`)
	fmt.Println(count)
	fmt.Println("--9. 字符串的重复")
	repeat := strings.Repeat("abc", 10)
	fmt.Println(repeat)

	fmt.Println("--10. 大小写转换")
	lower := strings.ToLower("ABC")
	fmt.Println(lower)
	upper := strings.ToUpper("abc")
	fmt.Println(upper)

	fmt.Println("--11. 拼接字符串")
	join := strings.Join([]string{"a", "b", "c"}, "-")
	fmt.Println(join)

	fmt.Println("--12. 字符串与其他类型之间的转换")
	size := strconv.IntSize
	fmt.Println("int size is :", size)
	itoa := strconv.Itoa(1)
	fmt.Println(itoa)
	atoi, err := strconv.Atoi("2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(atoi)
	}
}

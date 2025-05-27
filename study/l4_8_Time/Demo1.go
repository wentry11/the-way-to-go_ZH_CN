package main

import (
	fm "fmt"
	"time"
)

func main() {

	// 标准时间
	utc := time.Now().UTC()
	fm.Println(utc)
	// 当前时区时间
	fm.Println(time.Now())

	now := time.Now()
	week_ns := 60 * 60 * 24 * 7 * 1e9
	duration := time.Duration(week_ns)
	// 当前时间 加了一周
	fm.Println(now.Add(duration))

	// 按照时间格式进行格式化
	fm.Println(now.Format("2006-01-02 15:03:05"))
	// 这种方式不太行了
	fm.Println(now.Format("yyyy-MM-hh HH:mm:ss"))

	fm.Println(now.Format("20060201"))

	//
}

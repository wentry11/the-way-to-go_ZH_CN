package main

const name = 1
const (
	a = 1 + iota
	b
	c
	d = c + iota
)

func main() {
	println(a)
	println(b)
	println(c)
	println(d)
}

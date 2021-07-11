package main

import "fmt"

const (
	g = iota
	h
	i
)

/*
const(
	g = iota
	h = iota
	i = iota
)
这和上面的效果是一样的
*/

func main() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	const d, e = 10, 20
	fmt.Println(d, e)

	fmt.Println(g, h, i)
}

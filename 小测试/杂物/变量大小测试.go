package 杂物

import (
	"fmt"
	"reflect"
	"unsafe"
)

const (
	a int32= 10000
	b = 10.5
	c = "www"
	d = 1000000000000000000
	o
	p = true

)
var e = 10000
var f = 10.5
var g = "www"
var h = 1000000000000000000


var i = int64(10000)
var j = int32(10000)
var k = uint8(100)
func main() {
	show(a)
	show(b)
	show(c)
	show(d)
	show(o)
	show(p)

	show(e)
	show(f)
	show(g)
	show(h)

	show(i)
	show(j)
	show(k)
	/*
		int 16 8
		float64 16 8
		string 16 8
		int 16 8
		int 16 8
		float64 16 8
		string 16 8
		int 16 8
		int64 16 8
		int32 16 8
		uint8 16 8
	*/
}

func show(x interface{}) {
	fmt.Println(x,reflect.TypeOf(x),unsafe.Sizeof(x),unsafe.Sizeof(&x))
}

/*
	总结
	1. go语言普通变量是16个字节的，而指针变量是8个字节.. (结论有待商榷)
	2. 常量必须初始化

*/
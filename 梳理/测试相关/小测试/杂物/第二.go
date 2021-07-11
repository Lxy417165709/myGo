package 杂物

import (

)


var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {}
	print(a)
}
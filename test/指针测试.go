package main

import (
	"fmt"

)

type Node struct {

}
func function(a *Node){
	fmt.Println("函数体内",a,&a)
}


func main() {
	a:=&Node{}
	fmt.Println("函数体外",a,&a)
	function(a)

}
/*
	总结
	1. 函数体内的指针和函数体外的指针，指向的是同一个地方（值是一样的），
       但是他们的地址是不一样的，所以哈希时，直接哈希就可以了
       因为 mp[a]在函数体内外都是一样的，不用在函数体内写作mp[&(*a)],
	   mp[&(*a)] 其实也等于 mp[a]
*/
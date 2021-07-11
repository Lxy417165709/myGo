package main

import "fmt"

type I interface {
	Hello()
}
type S struct{

}
func (s *S) Hello(){
}


func main(){
	var i I = (*S)(nil)

	fmt.Println(i == nil,i)
	// false <nil>        // I 明明等于nil....为什么会是false呢？
}

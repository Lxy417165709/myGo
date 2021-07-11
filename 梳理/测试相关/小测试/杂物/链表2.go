package 杂物

import "fmt"

type Node struct {
	Val int
	Next *Node
}
func show(a *Node){
	for ;a!=nil ;  {
		fmt.Println(a.Val)
		a=a.Next
	}
}

func main() {
	a:=&Node{1,nil}
	a.Next = a
	a.Next.Val = 10
	show(a)
	// 自成环 0.0..
}

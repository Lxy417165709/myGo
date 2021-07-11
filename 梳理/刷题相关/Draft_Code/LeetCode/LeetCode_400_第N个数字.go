package main

import (
	"fmt"
)


type Node struct {
	l int
	r int
	time int
	l_num int
	r_num int
}

func findNthDigit(n int) int {
	return get(n)
}


func get(num int) int{
	l:=1
	r:=9
	slice := make([]Node,0)
	tmp :=0
	for i:=1;i<15 ;i++  {
		slice =append(slice, Node{tmp+1,tmp+(r-l+1)*i,i,l,r})
		tmp += (r-l+1)*i
		l = l*10
		r = l*10-1
	}
	index:=-1
	for i:=0;i< len(slice);i++  {
		if num>=slice[i].l && num<=slice[i].r {
			index = i
			break
		}
	}

	off := num-slice[index].l
	x:=off%slice[index].time
	y:=off/slice[index].time

	target := slice[index].l_num + y

	digits := make([]int,0)
	for ;target!=0 ;  {
		digits = append(digits,target%10 )
		target = target/10
	}

	return digits[len(digits)-1-x]





}


func main() {

	x:=0
	for ;true ;  {
		fmt.Scan(&x)
		fmt.Println(findNthDigit(x))
	}
}
/*
	总结
	1. 这题看起来复杂，但是一步一步来的话，其实会变得简单~
	2. 上面的代码易读性可能不太好，代码可以再优美一点的 0.0..

*/
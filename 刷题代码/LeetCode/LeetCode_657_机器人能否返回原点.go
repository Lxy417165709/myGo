package main

import "fmt"

func judgeCircle(moves string) bool {
	x,y:=0,0
	for i:=0;i< len(moves);i++  {
		if moves[i]=='U' {
			y++
		}
		if moves[i]=='D' {
			y--
		}
		if moves[i]=='R' {
			x++
		}
		if moves[i]=='L' {
			x--
		}
	}
	return (x==0)&&(y==0)
}
func main() {
	fmt.Println(judgeCircle(""))
}
/*
	总结
	1. 这是一道水题
*/
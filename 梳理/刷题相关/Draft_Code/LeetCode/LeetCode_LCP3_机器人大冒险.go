package main

import "fmt"

func min(a,b int) int{
	if	a>b{
		return b
	}
	return a
}
func robot(command string, obstacles [][]int, x int, y int) bool {
	xx,yy:=0,0

	m:=make(map[int]int)
	m[xx<<30 | yy] = 1
	for _,v:= range command  {
		if v=='U' {
			yy++
		}else {
			xx++
		}
		m[xx<<30 | yy] = 1
	}
	circle :=min(x/xx,y/yy)

	//x = x - circle*xx
	//y = y - circle*yy

	if m[(x - circle*xx)<<30|(y - circle*yy)]==0 {
		return false
	}
	for i:=0;i< len(obstacles);i++  {
		if obstacles[i][0]>x || obstacles[i][1]>y {
			continue
		}
		// 以下三行是去除循环，回归原始
		circle :=min(obstacles[i][0]/xx,obstacles[i][1]/yy)
		obstacles[i][0] = obstacles[i][0] - circle*xx
		obstacles[i][1] = obstacles[i][1] - circle*yy



		if m[(obstacles[i][0]<<30)|(obstacles[i][1])]==1 {
			return false
		}
	}
	return true

}


func main() {
	fmt.Println(robot("URR",[][]int{{0,0}},0,0))
}
/*
	总结
	1. 该题我是看官方的，然后自己写的
	2. 在做之前也有官方的想法，但是不知道怎么实现，也可以说是实现方式有点错误...
	3. 利用位运算可以很方便的进行哈希改造,如  m[x<<30)|y]
	4. 有循环的题目，可以通过去除循环，还原到最初的状态，比如点(x,y)，我们可以把它
	   去除到走完第一个循环的区域
*/
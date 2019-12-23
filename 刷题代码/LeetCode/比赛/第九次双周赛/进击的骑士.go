package 第九次双周赛

import (
	"fmt"
)


//func judge(x,y,border int) bool{
//	if x<0 || y<0 || x>border {
//		return false
//	}
//	return true
//}

//func set(x,y,number int){
//	dp[400+x][400+y] = number
//}
//
//func get(x,y int) int{
//	return dp[x+400][400+y]
//}
//var dp [800][800]int
//
//// 这dp有逻辑漏洞
//func minKnightMoves(x int, y int) int {
//	dp=[800][800]int{
//	}
//	for i:=0;i< len(dp);i++  {
//		for t:=0;t< len(dp[i]);t++  {
//			dp[i][t]=INF
//		}
//	}
//
//	set(0,0,0)
//	set(1,0,3)
//	set(1,1,3)
//	set(2,0,2)
//	set(2,1,1)
//	set(2,2,3)
//
//
//	A:=[][]int{
//		{-2,1},{-2,-1},{-1,2},{-1,-2},
//	}
//	for d:=3;d<=300 ;d++  {
//		for t:=0;t<=d ;t++  {
//			mn:=INF
//			for j:=0;j< len(A);j++  {
//				mn=min(mn,get(d+A[j][0],t+A[j][1]))
//			}
//			set(d,t,mn+1)
//		}
//	}
//
//	if x<0 {
//		x=-x
//	}
//	if y<0 {
//		y = -y
//	}
//	if y>x {
//		x,y = y,x
//	}
//
//
//
//	for i:=0;i<=5 ;i++  {
//		for t:=0;t<=5 ;t++  {
//			fmt.Print(get(i,t)," ")
//		}
//		fmt.Println()
//	}
//	return get(x,y)
//}
//func min(a,b int) int{
//	if a>b {
//		return b
//	}
//	return a
//}
const (
	INF = 1000000000
	OFF = 400
)
type Node struct {
	x int
	y int
}

func newNode(x,y int) Node{
	return Node{x,y}	// 只在标记和判断标记的时候+OFF，类似哈希 0.0.. 这不用加
}
func mark(x,y int){
	m[x+OFF][y+OFF] = 1
}
func judge(x,y int) bool{
	return !(m[x+OFF][y+OFF]==1)
}

var m [800][800]int
func minKnightMoves(x int, y int) int {
	m =[800][800]int{}
	dx:=[]int{-1,-1,-2,-2,1, 1,2, 2}
	dy:=[]int{-2, 2,-1, 1,2,-2,1,-1}


	queue:=make([]Node,0)
	queue = append(queue, Node{0,0})
	ans:=0
	for ; len(queue)!=0;  {
		ans++
		length:= len(queue)	// 注意这句，要保留当前queue的长度，不能在下面写i<len(queue)，因为在运行过程中，queue的长度会一直变..
		for i:=0;i<length ;i++  {
			top:=queue[0]
			queue = queue[1:]
			if top.x==x && top.y==y {
				return ans-1
			}
			for t:=0;t< len(dx);t++  {
				newX,newY :=top.x+dx[t],top.y+dy[t]
				if judge(newX,newY)==true {
					queue = append(queue, newNode(newX,newY))
					mark(newX,newY)
				}
			}
		}

	}

	return -1
}
func main() {
	fmt.Println(minKnightMoves(300,0))
}
/*
	总结
	1. 第一次使用BFS的时候，遍历方式错了
	2. 第二次觉得可以使用DP，但是发现DP的遍历方式有逻辑漏洞
	3. 第三次借鉴官方的BFS做 0.0.. 暴力BFS，上面有一些注意点
	4. 这次BFS的方式是自内到外拓展，之前写的那个是拓展8个位置，在到下个位置拓展8个位，然后就超时，也错误了..
*/
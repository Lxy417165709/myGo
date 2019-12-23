package main

import "fmt"

var mp [505][505]bool
var num int
func numEnclaves(A [][]int) int {
	ans:=0
	mp = [505][505]bool{}
	num=0
	n,m:=len(A),len(A[0])
	for i:=0;i< n;i++  {
		for t:=0;t< m;t++  {
			if A[i][t]==1 {
				ans++
			}
		}
	}


	for i:=0;i<n ;i++  {
		search(A,i,0)
		search(A,i,m-1)

	}
	for i:=0;i<m ;i++  {
		search(A,0,i)
		search(A,n-1,i)
	}

	return ans - num
}

func search(A [][]int,x int,y int){
	if x<0 || x>= len(A) || y<0 || y>= len(A[x]) {
		return
	}
	if mp[x][y]==false {
		mp[x][y]=true
		if A[x][y]==1 {
			num++
			search(A,x-1,y)
			search(A,x+1,y)
			search(A,x,y-1)
			search(A,x,y+1)
		}
	}
	return
}



func main() {
	//A:=[][]int{
	//	//	//	//	{0,0,0,0,0,0,0},
	//	//	//	//	{1,0,1,0,1,1,0},
	//	//	//	//	{0,1,1,0,1,1,0},
	//	//	//	//	{0,0,0,1,0,0,0},
	//	//	//	//}
	A:=[][]int{
		{0,1,1,0},
		{0,0,1,0},
		{0,0,1,0},
		{0,0,0,0},
	}


	fmt.Println(numEnclaves(A))


}

/*
	总结
	1. 该题是一道搜索题
	2. 注意初始化啊啊啊啊啊！！！！
*/
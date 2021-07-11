package main

import "fmt"
func max(a,b int) int{
	if a>b {
		return a

	}
	return b
}
var dx = []int{1,-1,0,0}
var dy = []int{0,0,1,-1}
func getMaximumGold(grid [][]int) int {
	ans:=0
	for i:=0;i< len(grid);i++  {
		for t:=0;t< len(grid[i]);t++  {
			m=make(map[int]int)
			x:=dfs(grid,i,t)
			fmt.Println(i,t,x)
			ans=max(ans,x)
		}
	}
	return ans

}
var m map[int]int
func dfs(grid [][]int,x int,y int) int{

	if x<0 ||x>= len(grid) || y<0 || y>=len(grid[0]) || m[x<<10|y]==1||grid[x][y]==0 {
		return 0
	}


	aa:=0
	ans:=0
	for i:=0;i< len(dx);i++  {
		m[x<<10|y]=1	// 很重要！
		aa=grid[x][y]+dfs(grid,x+dx[i],y+dy[i])
		m[x<<10|y]=0
		ans=max(aa,ans)
	}
	return ans

}


func main() {
	fmt.Println(getMaximumGold([][]int{
		{1,0,7,0,0,0},{2,0,6,0,1,0},{3,5,6,7,4,2},{4,3,1,0,2,0},{3,0,5,0,20,0},
	}))
}
/*
	总结
	1. 第一次错在没有清除标记..
*/
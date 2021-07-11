package main

import "fmt"

var mp map[[2]int] bool
var off uint8= 10
func closedIsland(grid [][]int) int {
	mp = make(map[[2]int] bool)
	ans:=0
	for i:=0;i<len(grid);i++{
		for t:=0;t<len(grid[i]);t++{
			if grid[i][t]== 1{
				continue
			}
			if _,ok:=mp[[2]int{i,t}];ok{
				continue
			}
			if dfs(i,t,grid)==true{
				ans++
			}
		}
	}
	return ans
}

func dfs(x,y int,grid [][]int) bool{
	if x<0 || y<0 || x>=len(grid)||y>=len(grid[0]){
		return false
	}
	// 记忆化搜索
	if a,ok:=mp[[2]int{x,y}];ok{
		return a
	}
	if grid[x][y]==1{
		return true
	}
	grid[x][y] = 1	// 标记
	mp[[2]int{x,y}]=dfs(x+1,y,grid) &&dfs(x-1,y,grid) &&dfs(x,y-1,grid) &&dfs(x,y+1,grid)
	return mp[[2]int{x,y}]

}

func main() {
	fmt.Println(closedIsland([][]int{
		{1,1,1,1,1,1,1,0},
		{1,0,0,0,0,1,1,0},
		{1,0,1,0,1,1,1,0},
		{1,0,0,0,0,1,0,1},
		{1,1,1,1,1,1,1,0},
	}))
}

/*
	总结
	1. 这题我使用了dfs, 对于每个陆地，我判断它的4个方位是否能到达海洋，如果可以则ans++
	2. 要注意访问后进行标记，否则会走回头路
*/
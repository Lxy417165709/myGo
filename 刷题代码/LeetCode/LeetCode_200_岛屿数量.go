package main

import "fmt"

var m [1050][1050]bool
func numIslands(grid [][]byte) int {
	m=[1050][1050]bool{}
	ans:=0
	for i:=0;i< len(grid);i++  {
		for t:=0;t< len(grid[i]);t++  {
			if grid[i][t]=='1' && m[i][t]==false {
				dfs(grid,i,t)
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]byte,i,j int){

	if i<0 || i>= len(grid){
		return
	}
	if j<0 || j>= len(grid[i]) {
		return
	}
	if grid[i][j] == '0' || m[i][j]==true {
		return
	}
	m[i][j] = true
	dfs(grid,i+1,j)
	dfs(grid,i-1,j)
	dfs(grid,i,j+1)
	dfs(grid,i,j-1)
	return
}


func main() {
	grid := [][]byte{
		{'1','1','0','1','1'},
		{'1','1','0','1','1'},
		{'1','1','0','1','1'},
		{'1','1','0','1','1'},
	}
	fmt.Println(numIslands(grid))
}
/*

	总结
	1. 这是一道简单的DFS题(当然，也可以用bfs做)
*/
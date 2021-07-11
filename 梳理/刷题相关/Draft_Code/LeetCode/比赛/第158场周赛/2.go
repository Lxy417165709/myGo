package main

import "fmt"

func abs(a int) int{
	if a>0{
		return a
	}
	return -a
}


var m map[int]int
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	m = make(map[int]int)
	m[king[0]<<10 + king[1]] = -1
	for i:=0;i< len(queens);i++  {
		m[queens[i][0]<<10 + queens[i][1]] = 1
	}
	ans:=make([][]int,0)
	for i:=0;i< len(queens);i++  {
		a:=dfs(queens[i][0]-1,queens[i][1],-1,0)
		b:=dfs(queens[i][0]+1,queens[i][1],1,0)
		c:=dfs(queens[i][0]-1,queens[i][1]-1,-1,-1)
		d:=dfs(queens[i][0]+1,queens[i][1]+1,1,1)
		e:=dfs(queens[i][0],queens[i][1]-1,0,-1)
		f:=dfs(queens[i][0],queens[i][1]+1,0,1)
		g:=dfs(queens[i][0]-1,queens[i][1]+1,-1,1)
		h:=dfs(queens[i][0]+1,queens[i][1]-1,1,-1)
		//fmt.Println(a,b,c,d,e,f,g,h)
		if a||b||c||d||e||f||g||h {
			ans = append(ans,queens[i])
		}
	}
	return ans
}
func dfs(x int,y int,dx int,dy int) bool{
	//fmt.Println(x,y)
	if x<0 || y<0 || y>=8 || x>=8 {
		return false
	}
	if m[x<<10 + y]==1 {
		return false
	}
	if m[x<<10 + y]==-1 {
		return true
	}
	return dfs(x+dx,y+dy,dx,dy)


}

func main() {
	fmt.Println(queensAttacktheKing([][]int{{5,6},{7,7},{2,1},{0,7},{1,6},{5,1},{3,7},{0,3},{4,0},{1,2},{6,3},{5,0},{0,4},{2,2},{1,1},{6,4},{5,4},{0,0},{2,6},{4,5},{5,2},{1,4},{7,5},{2,3},{0,5},{4,2},{1,0},{2,7},{0,1},{4,6},{6,1},{0,6},{4,3},{1,7}},[]int{
		3,4,
	}))
}

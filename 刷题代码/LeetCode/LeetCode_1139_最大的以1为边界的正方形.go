package main

import "fmt"


func max(a,b int) int{
	if a>b {
		return a

	}
	return b
}
func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}

func largest1BorderedSquare(grid [][]int) int {
	oneX:=make(map[int]int)
	oneY:=make(map[int]int)

	for i:=0;i< len(grid);i++  {
		for t:=0;t< len(grid[i]);t++  {
			if grid[i][t]==1 {
				oneX[i<<10+t] = oneX[i<<10+t-1]+1
				oneY[i<<10+t] = oneY[(i-1)<<10+t]+1
			}else{
				oneX[i<<10+t] = 0
				oneY[i<<10+t] = 0
			}
		}
	}
	ans:=0
	for i:=0;i< len(grid);i++  {
		for t:=0;t< len(grid[i]);t++  {

			if grid[i][t]==0 {
				continue
			}
			area:=min(len(grid[i])-t,len(grid)-i)
			for k:=area;k>ans ;k--  {
				ax,bx,cx,dx:=i,i,i+k-1,i+k-1
				ay,by,cy,dy:=t,t+k-1,t,k+t-1

				num1:=oneX[bx<<10+by] -oneX[ax<<10+ay-1]
				num2:=oneY[dx<<10+dy] -oneY[(bx-1)<<10+by]
				num3:=oneX[dx<<10+dy] -oneX[cx<<10+cy-1]
				num4:=oneY[cx<<10+cy] -oneY[(ax-1)<<10+ay]
				// num1!=0很关键，这表明不能组成正方形
				if num1==num2 && num2==num3 && num3==num4 && num1!=0 {
					ans = max(ans,num1)
					break
				}
			}

		}
	}
	return ans*ans

}

func main() {
	fmt.Println(largest1BorderedSquare([][]int{
		{1,0,},{0,0},
	}))
}
/*
	总结
	1. 这题和那个最大全1的正方形很像，但是时间复杂度不太一样..我是看了官方后用暴力做的
	2. 官方也有人用dp作出来的 0.0 可以学习,他的做法和我差不多 0.0..
	3. 这是优化过的版本，但是还268ms ...

*/
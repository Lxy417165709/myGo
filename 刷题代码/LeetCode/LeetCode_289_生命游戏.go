package main

import "fmt"

var off [8][2]int
var m [1050][1050]bool
func gameOfLife(board [][]int)  {
	off = [8][2]int{
		{-1,-1},
		{-1,0},
		{-1,1},
		{0,-1},
		{0,1},
		{1,-1},
		{1,0},
		{1,1},
	}
	m=[1050][1050]bool{}

	for i:=0;i< len(board);i++  {
		for t:=0;t< len(board[i]);t++  {
			m[i][t] = judge(board,i,t,len(board),len(board[i]))
		}
	}
	for i:=0;i< len(board);i++  {
		for t:=0;t< len(board[i]);t++  {
			if m[i][t]==false {
				if board[i][t]==0 {
					board[i][t]=1
				}else{
					board[i][t]=0
				}
			}
		}
	}
	fmt.Println(board)
}

func judge(board [][]int,x,y,n,m int) bool{

	xb:=0
	hxb:=0
	for i:=0;i< len(off);i++  {
		xx:=x+off[i][0]
		yy:=y+off[i][1]
		if xx>=0 && xx< n && yy>=0 && yy<m{
			xb++
			if board[xx][yy]==1 {
				hxb++
			}
		}
	}
	if board[x][y]==1 {
		if hxb==2 || hxb==3 {
			return true
		}else{
			return false
		}
	}else{
		if hxb==3 {
			return false
		}else{
			return true
		}
	}
}




func main() {
	gameOfLife([][]int{
		{1,1,1},
		{1,0,1},
		{1,0,1},
		{1,1,1},
	})
}
/*
	总结
	1. 这是一道简单题~通过judge遍历标记，如果是false则翻转其值(0->1,1->0)
*/
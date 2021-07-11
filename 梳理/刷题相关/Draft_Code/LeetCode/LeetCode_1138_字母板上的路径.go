package main

import "fmt"

type path struct{
	x,y int
	s string
}
func alphabetBoardPath(target string) string {

	beginX,beginY:=0,0
	ans:=""
	for i:=0;i< len(target);i++  {
		path:=bfs(beginX,beginY,target[i])
		fmt.Println(string(target[i]),path)
		beginX,beginY = path.x,path.y
		ans=ans + path.s + "!"
	}
	return ans
}
func bfs(x,y int,target uint8) path{
	m:=make(map[int]int)
	queue:=make([]path,0)

	board:=[]string{
		"abcde","fghij","klmno","pqrst","uvwxy","z",
	}
	dx:=[]int{0,0,-1,1}
	dy:=[]int{1,-1,0,0}
	firth:=path{x,y,""}
	queue = append(queue, firth)
	m[100*x+y] = 1	// 因为这没改100*x+y啊啊啊啊啊
	for len(queue)!=0{
		top:=queue[0]
		queue=queue[1:]
		//fmt.Println(top)
		if board[top.x][top.y]==target {
			return top
		}

		for i:=0;i< len(dx);i++  {
			newX := dx[i]+top.x
			newY := dy[i]+top.y
			//if top.x==4 && top.y==0 {
			//	fmt.Println(newX,newY,m[5*newY+newX],"------")
			//}
			if newX<0 || newY<0 || newY>4 || newX>5 || newY+5*newX>25|| m[100*newX+newY] == 1  {
				continue
			}
			direct:=""
			switch i {
			case 0:
				direct="R"
			case 1:
				direct="L"
			case 2:
				direct="U"
			case 3:
				direct="D"
			}

			queue = append(queue, path{newX,newY,top.s + direct})
			m[100*newX+newY] = 1
		}
	}
	return path{-1,-1,""}
}




func main() {
	fmt.Println(alphabetBoardPath("kcz"))
}
/*
	总结
	1. 第一次错是因为哈希不正确 ,newY*5 + newX 会导致 (5,0)这个点和(1,0)这个点的哈希值一样 0.0. 所以我改用更大的 newX*100+newY
	2. 第二次错是因为 BFS中5*y+x改100*x+y后，但是入队开始时没改..
	3. 我是用BFS做的，但其实还有更简单的方法，官方有0.0.
*/
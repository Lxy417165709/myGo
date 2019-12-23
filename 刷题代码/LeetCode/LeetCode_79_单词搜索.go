package main



func exist(board [][]byte, word string) bool {
	if len(board)==0{
		return false
	}
	mark = make(map[int]int)
	m,n = len(board),len(board[0])

	for i:=0;i<m;i++{
		for t:=0;t<n;t++{
			// mark[i<<off|t] = 1  //这一步很重要，进去的时候要标记
			flag:=solve(board,i,t,word)
			// mark[i<<off|t] = 0  //这一步很重要，出来的时候要解除标记
			if flag==true{
				return true
			}
		}
	}
	return false

}

var mark map[int]int
var m,n int
var dx = []int{0,0,-1,1}
var dy = []int{1,-1,0,0}
const off = 20
func solve(board [][]byte,x int,y int,word string) bool{

	if !judge(x,y){
		return false
	}
	mark[x<<off|y] = 1
	if word[0]!=board[x][y]{
		mark[x<<off|y] = 0
		return false
	}

	word = word[1:]
	if word==""{
		mark[x<<off|y] = 0
		return true
	}


	for i:=0;i<len(dx);i++{
		newX,newY := x+dx[i],y+dy[i]
		if judge(newX,newY)==false {
			continue
		}
		//mark[newX<<off|newY] = 1
		flag:=solve(board,newX,newY,word)
		//mark[newX<<off|newY] = 0

		if flag==true{
			mark[x<<off|y] = 0
			return true
		}
	}
	mark[x<<off|y] = 0
	return false
}

func judge(x,y int) bool{
	if x>=0 && x<=m-1 && y>=0 && y<=n-1 && mark[x<<off|y]==0{
		return true
	}
	return  false
}

func main() {

}

/*
	总结
	1. DFS现在我越来越熟练了，但是写的还是存在一些小BUG 0.0
	2. 要注意标记！！注意点在上面了！
	3. 之后我修改了一些，更加统一了对点的标记
*/

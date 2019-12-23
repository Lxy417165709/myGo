package main



func numEnclaves(A [][]int) int {
	ans:=0
	m,n:=len(A),len(A[0])
	for i:=0;i<m;i++{
		solve(A,i,0)
		solve(A,i,n-1)
	}
	for i:=0;i<n;i++{
		solve(A,0,i)
		solve(A,m-1,i)
	}
	for i:=0;i<m;i++{
		for t:=0;t<n;t++{
			ans+= A[i][t]
		}
	}

	return ans
}

func solve(A [][]int,x int,y int){
	if x<0 || x>=len(A) || y<0 || y>=len(A[0]){
		return
	}
	if A[x][y]==0{
		return
	}

	A[x][y] = 0
	solve(A,x+1,y)
	solve(A,x,y+1)
	solve(A,x-1,y)
	solve(A,x,y-1)
}


/*
	总结
	1. 这方法是扫描边界，把边界陆地填充为海洋，之后再计算陆地数量。
*/
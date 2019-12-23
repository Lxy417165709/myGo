package main

func minMutation(start string, end string, bank []string) int {
	mark = make(map[int]int)
	flag:=0
	for i:=0;i<len(bank);i++{
		if end==bank[i]{
			flag = 1
			break
		}
	}
	if flag==0{
		return -1
	}
	ans:=solve(start,end,bank)
	if ans==100000000{
		return -1
	}else{
		return ans
	}
}

var mark map[int]int
func solve(start string,end string,bank []string) int{

	x:=test(start,end)
	if x<=1{
		for i:=0;i<len(bank);i++{
			if end==bank[i]{
				return x
			}
		}
		return 100000000
	}


	ans:=100000000
	for i:=0;i<len(bank);i++{
		if mark[i]==1{
			continue
		}
		y:=test(bank[i],end)

		if y<=1{
			mark[i]=1
			ans = min(solve(start,bank[i],bank)+y,ans)
			mark[i]=0
		}
	}
	return ans
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}
func test(a,b string) int{
	x:=0
	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			x++
		}
	}
	return x
}

func main() {

}
/*
	总结
	1. 这题我是使用深度优先搜索AC的，但是官方题解大多使用bfs 0.0
	2. 我也觉得bfs好，比较省时间
*/
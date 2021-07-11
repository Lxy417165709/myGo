package main
func solveNQueens(n int) [][]string {
	ans = make([][]string,0)
	mp = make(map[int]int)
	bytes := make([]byte,n)
	for i:=0;i<n;i++{
		bytes[i] = '.'
	}
	blockString = string(bytes)
	solve(n,0,[]string{})
	return ans
}

var mp map[int]int
var ans [][]string
var blockString string
func solve(n int,lay int,sli []string){
	if lay==n{
		tmp:=make([]string,len(sli))
		for i:=0;i<len(tmp);i++{
			tmp[i] = sli[i]
		}
		ans = append(ans,tmp)
		return
	}


	for i:=0;i<n;i++{
		if mp[i<<10|lay]==0{
			set(i,lay,n,1)
			sli = append(sli,getQstring(i))
			solve(n,lay+1,sli)
			sli = sli[:len(sli)-1]
			set(i,lay,n,-1)
		}
	}
}

func getQstring(x int) string{
	bytes:=[]byte(blockString)
	bytes[x] = 'Q'
	return string(bytes)
}

func set(x int,lay int,n int,num int){
	for i:=0;i<n;i++{
		mp[(x+i)<<10|(i+lay)] += num
		mp[x<<10|(i+lay)] += num
		mp[(x-i)<<10|(i+lay)] += num
	}
}



func main() {

}


/*
	总结
	1. 该题我使用了搜索的算法，搜索到第n+1层表示有解，此时就加入结果集
	2. 第一次mp标记的时候采用了0 1，表示该位置没有or已有棋子，但是这样会导致回溯
	   清除修改的时候，把上层的标记也清除了，于是我采用了数字，mp = 0 1 2，表示该
		位置被0,1,2,棋子占据（不一定是在该位置上）
	3. 官方的话是标记行列加点的，那个可以剩一些时间，我这个是把棋盘都看作是点的
	   如果不允许放，则mp值！=0
*/
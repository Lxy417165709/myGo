package main


func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	sum:=0
	for i:=0;i<len(colsum);i++{
		sum+=colsum[i]
	}
	if upper+lower!=sum{
		return [][]int{}
	}
	up:=make([]int,len(colsum))
	low:=make([]int,len(colsum))
	for i:=0;i<len(colsum) ;i++  {

		if upper <0 || lower<0 {
			return [][]int{}
		}

		if colsum[i]==2 {
			up[i]=1
			low[i]=1
		}
		if colsum[i]==0 {
			up[i]=0
			low[i]=0
		}
		if colsum[i]==1 {
			if upper>lower{
				upper--
				up[i]=1
			}else{
				lower--
				low[i]=1
			}
		}
	}

	return [][]int{up,low}
}
/*
func solve(upper int, lower int, colsum []int) [][]int{
	if upper<0 || lower<0{
		return [][]int{}
	}
	if len(colsum)==1{
		if upper<=1 && lower<=1 && colsum[0]==upper+lower{
			ans:=make([][]int,0)
			ans = append(ans,[]int{upper})
			ans = append(ans,[]int{lower})
			return ans
		}
		return [][]int{}
	}



	if colsum[0]==1{
		up:=solve(upper-1,lower,colsum[1:])
		low:=solve(upper,lower-1,colsum[1:])
		ans:=make([][]int,2)
		if len(up)!=0{
			ans[0] = append(ans[0],1)
			ans[1] = append(ans[1],0)
			for i:=0;i<len(colsum)-1;i++{
				ans[0] = append(ans[0],up[0][i])
				ans[1] = append(ans[1],up[1][i])
			}
			return ans
		}
		if len(low)!=0{
			ans[0] = append(ans[0],0)
			ans[1] = append(ans[1],1)
			for i:=0;i<len(colsum)-1;i++{
				ans[0] = append(ans[0],low[0][i])
				ans[1] = append(ans[1],low[1][i])
			}
			return ans
		}

	}
	if colsum[0]==0{
		x:=solve(upper,lower,colsum[1:])
		ans:=make([][]int,2)
		if len(x)!=0{
			ans[0] = append(ans[0],0)
			ans[1] = append(ans[1],0)
			for i:=0;i<len(colsum)-1;i++{
				ans[0] = append(ans[0],x[0][i])
				ans[1] = append(ans[1],x[1][i])
			}
			return ans
		}
	}
	if colsum[0]==2{
		x:=solve(upper-1,lower-1,colsum[1:])
		ans:=make([][]int,2)
		if len(x)!=0{
			ans[0] = append(ans[0],1)
			ans[1] = append(ans[1],1)
			for i:=0;i<len(colsum)-1;i++{
				ans[0] = append(ans[0],x[0][i])
				ans[1] = append(ans[1],x[1][i])
			}
			return ans
		}
	}
	return [][]int{}
}
*/

func main() {

}

/*
	总结
	1. 第一个使用动态规划，然后发现时间不符合题意，之后看了看，才发现了贪心性质
		- 如果colomsum[i]==2 则up[i],low[i]都要为1,之后upper--,lower--就可以了
		- 如果colomsum[i]==1 则选择upper和lower之中最大的标1，并做相应的减少
		- 如果colomsum[i]==0 则up[i],low[i]都要为0
		- 循环完毕，如果upper!=0 || lower!=0 则不满足题意，返回[]
	2. 不知道为什么，用go提交一直超时，用c++直接AC
*/
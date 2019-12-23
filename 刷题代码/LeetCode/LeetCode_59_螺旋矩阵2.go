package main


func generateMatrix(n int) [][]int {
	return solve(n,1)
}

// 递归写法
//func solve(n int,number int) [][]int{
//
//	ans := make([][]int,0)
//
//	for i:=0;i<n;i++{
//		ans = append(ans,make([]int,n))
//	}
//	if n==0{
//		return ans
//	}
//	if n==1{
//		ans[0][0] = number
//		return ans
//	}
//
//	beginX,beginY:=0,0
//	endX,endY:=n-1,n-1
//
//	for t:=0;t<endY;t++{
//		ans[0][t] = number
//		number++
//	}
//	for i:=0;i<endX;i++{
//		ans[i][endY] = number
//		number++
//	}
//	for t:=endY;t>0;t--{
//		ans[endX][t] = number
//		number++
//	}
//
//	for i:=endX;i>0;i--{
//		ans[i][0] = number
//		number++
//	}
//	next:=solve(n-2,number)
//
//	// 合并
//	for i:=0;i<len(next);i++{
//		for t:=0;t<len(next[i]);t++{
//			ans[i+beginX+1][t+beginY+1] = next[i][t]
//		}
//	}
//
//	return ans
//
//}

// 迭代写法
func solve(n int,number int) [][]int{

	ans := make([][]int,0)

	for i:=0;i<n;i++{
		ans = append(ans,make([]int,n))
	}


	beginX,beginY:=0,0
	endX,endY:=n-1,n-1
	number = 1
	for number<=n*n{
		for t:=beginY;t<=endY;t++{
			ans[beginX][t] = number
			number++
		}
		beginX++
		for i:=beginX;i<=endX;i++{
			ans[i][endY] = number
			number++
		}
		endY--
		for t:=endY;t>=beginY;t--{
			ans[endX][t] = number
			number++
		}
		endX--

		for i:=endX;i>=beginX;i--{
			ans[i][beginY] = number
			number++
		}
		beginY++
	}

	return ans

}
func main() {

}

/*
	总结
	1. 这题我是使用剥皮法 + 递归完成的, 每一次递归剥一层皮，
		则一层皮剥完成后，下一次的就是solve(n-2)，之后再把结果合并起来 0.0
	2. 我把迭代写法也写了一次
*/

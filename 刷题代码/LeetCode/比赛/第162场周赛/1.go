package main


func oddCells(n int, m int, indices [][]int) int {

	ans :=[100][100]int{}

	for i:=0;i<len(indices);i++{
		row:=indices[i][0]
		col:=indices[i][1]
		for t:=0;t<m;t++{
			ans[row][t]++
		}
		for t:=0;t<n;t++{
			ans[t][col]++
		}
	}
	cnt:=0
	for i:=0;i<n;i++{
		for t:=0;t<m;t++{
			cnt += ans[i][t]&1
		}
	}
	return cnt


}

func main() {

}

/*
	总结
	1. 这个题目只需要模拟就可以了
*/
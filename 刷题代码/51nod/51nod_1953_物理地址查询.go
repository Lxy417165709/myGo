package _1nod

import (
	"fmt"
	"strconv"
)

const INF = 2<<60 - 1

var m [105][105]int64
var dis [105]int64
var mark [105]int

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		dis[i] = INF
		mark[i] = 0
	}
	for i := 2; i <= n; i++ {
		for t := 1; t <= i-1; t++ {
			var s string
			fmt.Scan(&s)
			if s == "x" {
				m[i][t] = INF
				m[t][i] = INF
			} else {
				x, _ := strconv.Atoi(s)
				m[i][t] = int64(x)
				m[t][i] = int64(x)
			}
		}
	}

	//for i:=1;i<=n;i++{
	//	for t:=1;t<=n ;t++  {
	//		fmt.Print(m[i][t]," ")
	//	}
	//	fmt.Println()
	//}
	dis[1] = 0

	for i := 1; i <= n; i++ {

		u := 1
		var minDis int64 = INF
		for t := 1; t <= n; t++ {
			if mark[t] == 1 {
				continue
			}
			if minDis > dis[t] {
				u = t
				minDis = dis[t]
			}
		}
		mark[u] = 1
		for t := 1; t <= n; t++ {
			if dis[t] > dis[u]+m[u][t] {
				dis[t] = dis[u] + m[u][t]
			}
		}
	}
	var time int64 = 0
	for k := 1; k <= n; k++ {
		if time<dis[k]{
			time = dis[k]
		}
	}
	fmt.Println(time)
}

/*
	知识点
	1. 没想出,但是迪杰斯特拉算法是我自己手打的(没看书)!!!必须表扬!!!

*/
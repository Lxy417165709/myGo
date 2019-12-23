package main

import "fmt"

type edge struct {
	to string
	value float64
}

const (
	INF = 1e30
)

var m map[string] []*edge
var mp map[string] int
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m = make(map[string] []*edge)
	for i:=0;i< len(equations);i++  {
		m[equations[i][0]] = append(m[equations[i][0]], &edge{
			equations[i][1],
			values[i],
		})
		x:=-1.0
		if values[i]==0 {
			x=INF
		}else{
			x=1.0/values[i]
		}
		m[equations[i][1]] = append(m[equations[i][1]], &edge{
			equations[i][0],
			x,
		})
	}
	ans:=make([]float64,0)
	for i:=0;i< len(queries);i++  {
		mp = make(map[string] int)
		if _,ok:=m[queries[i][0]];!ok {
			ans = append(ans, -1.0)
			continue
		}

		a,b:=solve(queries[i][0],queries[i][1])
		if b==true {
			ans = append(ans, a)
		}else{
			ans = append(ans, -1.0)
		}
	}
	return ans
}
// 这2个返回值其实可以变为float64就可以了，因为题目限制了正常返回值的范围 它指出结果值均为正数
func solve(first string,finaly string) (float64,bool){
	if first==finaly {
		return 1.0,true
	}
	mp[first] = 1
	edges:=m[first]
	for i:=0;i< len(edges);i++  {
		if mp[edges[i].to]==0 {
			mp[edges[i].to] = 1
			x,flag:=solve(edges[i].to,finaly)
			if flag==true {
				x=x*edges[i].value
				return x,true
			}
		}
	}
	return -1.0,false
}


func main() {
	fmt.Println(calcEquation([][]string{
		{"A","B"},{"B","C"},
	},[]float64{
		0,1,
	},[][]string{
		{"A","A"},{"A","C"},
	}))
}

/*
	总结
	1. 该题目我使用了深搜的思维
	2. 由于题目说明了除数不为0而且结果都为整数(除了非法的-1外)，所以其实可以省略好多判断条件的
	3. 官方题解还有用并查集的 0.0..
*/
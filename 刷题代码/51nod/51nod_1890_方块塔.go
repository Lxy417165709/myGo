package _1nod

import (
	"fmt"
	"sort"
)

type Box struct {

	l int
	w int
}
var m [5005]Box
var mark [5005]int
type boxSli []Box
type boxSli2 []Box
func (x boxSli) Len() int{
	return len(x)
}
func (x boxSli) Swap(i,j int){
	x[i],x[j] = x[j],x[i]
}
func (x boxSli) Less(i,j int) bool{
	if x[i].l != x[j].l{
		return x[i].l < x[j].l
	}else{
		return x[i].w < x[j].w
	}
}

func (x boxSli2) Len() int{
	return len(x)
}
func (x boxSli2) Swap(i,j int){
	x[i],x[j] = x[j],x[i]
}
func (x boxSli2) Less(i,j int) bool{
	if x[i].w != x[j].w{
		return x[i].w < x[j].w
	}else{
		return x[i].l < x[j].l
	}
}
/*
func main() {
	var n int
	fmt.Scan(&n)
	for i:=0;i<n ;i++  {
		fmt.Scan(&m[i].l)
		fmt.Scan(&m[i].w)
	}
	sort.Sort(boxSli(m[:n]))
	fmt.Println(m[:n])
	cnt:=1
	for i:=0;i<n-1 ;i++  {
		if m[i].w > m[i+1].w {
			cnt++
		}
	}
	sort.Sort(boxSli2(m[:n]))
	//fmt.Println(m[:n])
	cnt2:=1
	for i:=0;i<n-1 ;i++  {
		if m[i].l > m[i+1].l {
			cnt2++
		}
	}

	if cnt2 <cnt{
		cnt = cnt2
	}

	fmt.Println(cnt)
}
*/

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i].l)
		fmt.Scan(&m[i].w)
	}
	sort.Sort(boxSli(m[:n]))
	cnt:=0
	for i := 0; i < n; i++ {
		if mark[i]== 1 {
			continue
		}
		mark[i] = 1
		w:= m[i].w
		for t:=i+1;t<n;t++{
			if mark[t]==0{
				if m[t].w >= w{
					w = m[t].w
					mark[t] = 1
				}
			}
		}
		cnt++
	}
	fmt.Println(cnt)
}



/*
	总结
	1. 错误的做法是按L或者W排序后，判断是不是(如果排L，则判断W)递增的，这种方法
		对于{361 3599} {453 369} {544 5134}是错误的，因为满足条件的不一定紧跟其后

*/